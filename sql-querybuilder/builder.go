package sqlquerybuilder

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

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

// === SQL Table ===
type SQLTable struct {
	builder   *SQLQueryBuilder
	tableName string
	columns   map[string]string
}

func (builder *SQLQueryBuilder) NewTable(tableName string) *SQLTable {
	table := &SQLTable{
		builder:   builder,
		tableName: tableName,
		columns:   make(map[string]string),
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

func (table *SQLTable) DefineColumn(colName string, colType string) {
	table.columns[colName] = colType
}
func (table *SQLTable) EnsureTableExistsInDB() {
	colDef := ""
	for k, v := range table.columns {
		if colDef == "" {
			colDef = fmt.Sprintf("%s %s", k, v)
		} else {
			colDef = fmt.Sprintf("%s, %s %s", colDef, k, v)
		}
	}
	defineTable := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (\n%s\n);", table.tableName, colDef)
	fmt.Println(defineTable)
	//table.builder.db_conn.Exec(defineTable)
}

// === SQL Query Methods ===
type SQLQuery struct {
	table          *SQLTable
	whereCondition string
}

// Start a new query statement
func (table *SQLTable) NewQuery() *SQLQuery {
	query := &SQLQuery{
		table: table,
	}
	return query
}

func (query *SQLQuery) Where(condition string) *SQLQuery {
	query.whereCondition = condition
	return query
}

func (query *SQLQuery) FindAll() {
	q := fmt.Sprintf("SELECT * FROM %s WHERE %s;", query.table.tableName, query.whereCondition)
	fmt.Println(q)
}
func (query *SQLQuery) FindOne(out any) {
	q := fmt.Sprintf("SELECT * FROM %s WHERE %s;", query.table.tableName, query.whereCondition)
	fmt.Println(q)
	//query.table.builder.db_conn.Query("")
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

func (insert *SQLInsert) AddIntColumn(column string, val int) {
	insert.columns[column] = fmt.Sprintf("%d", val)
}
func (insert *SQLInsert) AddStringColumn(column string, val string) {
	insert.columns[column] = fmt.Sprintf("\"%s\"", val)
}

// Send the sql insert statement to the database
func (insert *SQLInsert) Send() {
	colDef := ""
	colVal := ""
	for k, v := range insert.table.columns {
		if colDef == "" {
			colDef = k
		} else {
			colDef = fmt.Sprintf("%s, %s", colDef, k)
		}
		val := insert.columns[k]
		if val == "" {
			if strings.Contains(v, "VARCHAR") {
				val = "\"\""
			} else {
				val = " "
			}
		}
		if colVal == "" {
			colVal = val
		} else {
			colVal = fmt.Sprintf("%s, %s", colVal, val)
		}
	}
	q := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s);", insert.table.tableName, colDef, colVal)
	fmt.Println(q)
	//insert.table.builder.db_conn.Exec("")
}
