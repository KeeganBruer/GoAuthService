package sqlquerybuilder

import (
	"database/sql"
	"fmt"
	"log"
)

type SQLQueryBuilder struct {
	db_conn *sql.DB
}
type SQLTable struct {
	builder   *SQLQueryBuilder
	tableName string
	columns   map[string]string
}
type SQLQuery struct {
	table          *SQLTable
	whereCondition string
}
type SQLInsert struct {
	table *SQLTable
}

func NewSQLQueryBuilder() *SQLQueryBuilder {
	builder := &SQLQueryBuilder{}
	return builder
}

func (builder *SQLQueryBuilder) Connect() {
	db, err := sql.Open("mysql",
		"user:password@tcp(127.0.0.1:3306)/hello")
	if err != nil {
		log.Fatal(err)
	}
	builder.db_conn = db
}

func (builder *SQLQueryBuilder) GetTable(tableName string) *SQLTable {
	table := &SQLTable{
		builder:   builder,
		tableName: tableName,
		columns:   make(map[string]string),
	}
	return table
}

// === SQL Query Methods ===

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
	q := fmt.Sprintf("SELECT * FROM %s WHERE %s", query.table.tableName, query.whereCondition)
	fmt.Println(q)
}
func (query *SQLQuery) FindOne(out any) {
	q := fmt.Sprintf("SELECT * FROM %s WHERE %s", query.table.tableName, query.whereCondition)
	fmt.Println(q)
	//query.table.builder.db_conn.Query("")
}

// === SQL Insert Methods ===

// Start a new insert statement
func (table *SQLTable) NewInsert() *SQLInsert {
	insert := &SQLInsert{
		table: table,
	}
	return insert
}

// Send the sql insert statement to the database
func (insert *SQLInsert) Send() {
	q := fmt.Sprintf("INSERT INTO %s (%s) COLUMNS ()", insert.table.tableName, "")
	fmt.Println(q)
	insert.table.builder.db_conn.Exec("")
}
