package sqlquerybuilder

import (
	"database/sql"
	"fmt"
	"log"
	"slices"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
)

type SQLQueryBuilder struct {
	db_conn *sql.DB
	tables  map[string]*SQLTable
}

func NewSQLQueryBuilder() *SQLQueryBuilder {
	builder := &SQLQueryBuilder{
		tables: make(map[string]*SQLTable),
	}
	return builder
}

func (builder *SQLQueryBuilder) Connect(cfg *mysql.Config) {
	// Get a database handle.
	var db *sql.DB
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
	builder.db_conn = db

}
func (builder *SQLQueryBuilder) UseDatabase(dbName string) {
	stm := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s;", dbName)
	fmt.Println(stm)
	_, err := builder.db_conn.Exec(stm)
	if err != nil {
		fmt.Println(err)
	}

	useStm := fmt.Sprintf("USE %s;", dbName)
	fmt.Println(useStm)
	_, err = builder.db_conn.Exec(useStm)
	if err != nil {
		fmt.Println(err)
	}
}

// === SQL Table ===
type SQLTable struct {
	builder     *SQLQueryBuilder
	tableName   string
	columns     map[string]string
	columnOrder []string
}

func (builder *SQLQueryBuilder) NewTable(tableName string) *SQLTable {
	table := &SQLTable{
		builder:     builder,
		tableName:   tableName,
		columns:     make(map[string]string),
		columnOrder: make([]string, 0),
	}
	return table
}
func (builder *SQLQueryBuilder) DefineTable(table *SQLTable) {
	builder.tables[table.tableName] = table
}
func (builder *SQLQueryBuilder) GetTable(tableName string) *SQLTable {
	table := builder.tables[tableName]
	return table
}
func (builder *SQLQueryBuilder) Int2DB(num int) string {
	return fmt.Sprintf("%d", num)
}
func (builder *SQLQueryBuilder) String2DB(str string) string {
	return fmt.Sprintf("\"%s\"", str)
}
func (builder *SQLQueryBuilder) Date2DB(timeStmp time.Time) string {
	dateTime := timeStmp.Format(time.DateTime)
	return fmt.Sprintf("\"%s\"", dateTime)
}

func (table *SQLTable) DefineColumn(colName string, colType string) {
	table.columns[colName] = colType
	if !slices.Contains(table.columnOrder, colName) {
		table.columnOrder = append(table.columnOrder, colName)
	}
}
func (table *SQLTable) EnsureTableExistsInDB() {
	colDef := ""
	for i := range table.columnOrder {
		k := table.columnOrder[i]
		v := table.columns[k]
		if colDef == "" {
			colDef = fmt.Sprintf("%s %s", k, v)
		} else {
			colDef = fmt.Sprintf("%s, %s %s", colDef, k, v)
		}
	}
	defineTable := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (\n%s\n);", table.tableName, colDef)
	fmt.Println(defineTable)
	table.builder.db_conn.Exec(defineTable)
}

// === SQL Query Methods ===
type SQLQuery struct {
	table          *SQLTable
	whereCondition string
}

// Start a new query statement
func (table *SQLTable) NewSelect() *SQLQuery {
	query := &SQLQuery{
		table: table,
	}
	return query
}

func (query *SQLQuery) Where(condition string) *SQLQuery {
	query.whereCondition = condition
	return query
}
func (query *SQLQuery) GetStatement() string {
	stm := fmt.Sprintf("SELECT %s", "*")
	stm = fmt.Sprintf("%s FROM %s", stm, query.table.tableName)
	if query.whereCondition != "" {
		stm = fmt.Sprintf("%s WHERE %s", stm, query.whereCondition)
	}
	stm = stm + ";"
	return stm
}

func (query *SQLQuery) FindAll(EachCB func(get func(dest ...any) error) error) error {
	statement := query.GetStatement()
	fmt.Println(statement)
	rows, err := query.table.builder.db_conn.Query(statement)
	if err != nil {
		return fmt.Errorf("query error %v", err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		err := EachCB(func(dest ...any) error {
			if err := rows.Scan(dest...); err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return fmt.Errorf("albumsByArtist %v", err)
		}
	}
	if err := rows.Err(); err != nil {
		return fmt.Errorf("albumsByArtist %v", err)
	}
	return nil
}
func (query *SQLQuery) FindOne(dest ...any) error {
	statement := query.GetStatement()
	fmt.Println(statement)

	row := query.table.builder.db_conn.QueryRow(statement)
	if err := row.Scan(dest...); err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("albumsById: no such album")
		}
		return fmt.Errorf("albumsById: %v", err)
	}
	return nil
}
func (query *SQLQuery) Exists() bool {
	statement := query.GetStatement()
	fmt.Println("Checking Exists:", statement)
	row := query.table.builder.db_conn.QueryRow(statement)
	if err := row.Scan(nil); err != nil {
		return false
	}
	return true
}

// === SQL Insert Methods ===
type SQLInsert struct {
	table   *SQLTable
	columns map[string]string
}

// Start a new insert statement
func (table *SQLTable) NewInsert() *SQLInsert {
	insert := &SQLInsert{
		table:   table,
		columns: make(map[string]string),
	}
	return insert
}

func (insert *SQLInsert) AddColumn(column string, val string) {
	insert.columns[column] = val
}

func (insert *SQLInsert) GetStatement() string {
	colDef := ""
	colVal := ""
	for i := range insert.table.columnOrder {
		k := insert.table.columnOrder[i]
		v := insert.table.columns[k]
		//If first item
		if colDef == "" {
			colDef = k
		} else {
			//all other items are appended with a proceeding comma
			colDef = fmt.Sprintf("%s, %s", colDef, k)
		}

		//process the data-to-be-inserted's matching value
		val := insert.columns[k]
		if val == "" { //it the value is empty
			if strings.Contains(v, "VARCHAR") { //if the table column def is a string
				//empty string column
				val = "\"\""
			} else {
				//null value
				val = " "
			}
		}
		if colVal == "" { //first value added
			colVal = val
		} else {
			//all other items are appended with a proceeding comma, like the coldefs
			colVal = fmt.Sprintf("%s, %s", colVal, val)
		}
	}
	//build statement
	stm := fmt.Sprintf("REPLACE INTO %s", insert.table.tableName)
	stm = fmt.Sprintf("%s (%s) VALUES (%s)", stm, colDef, colVal)

	stm = stm + ";"
	return stm
}

// Send the sql insert statement to the database
func (insert *SQLInsert) Send() bool {
	statement := insert.GetStatement()
	fmt.Println(statement)
	results, err := insert.table.builder.db_conn.Exec(statement)
	if err != nil {
		fmt.Println(err)
		return false
	}
	rows, err := results.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return false
	}
	return rows > 0
}
