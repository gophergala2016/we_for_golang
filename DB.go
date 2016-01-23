package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *MySqlDbClient

type MySqlDbClient struct {
	config *ConfigParams
	sqlDb  *sql.DB
}

func initializeDB(configparams ConfigParams) {
	db = NewMySqlDbClient(&configparams)
}

func NewMySqlDbClient(config *ConfigParams) *MySqlDbClient {
	return &MySqlDbClient{config: config}
}

func (db *MySqlDbClient) Open() error {
	var err error
	connectionString := fmt.Sprintf("%s:%s@/%s", db.config.user, db.config.password, db.config.database)
	if db.sqlDb, err = sql.Open("mysql", connectionString); err != nil {
		return err
	}
	if err = db.sqlDb.Ping(); err != nil {
		return err
	}
	return nil
}

func (db *MySqlDbClient) Close() {
	db.sqlDb.Close()
}

func (db *MySqlDbClient) ShowTables() {
	rows, err := db.sqlDb.Query("SHOW TABLES")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			panic(err)
		}
		log.Printf("%s\n", name)
	}
	if err := rows.Err(); err != nil {
		log.Println(err)
	}
}

func (db *MySqlDbClient) GenerateSchemaFile() error {

	var (
		rows *sql.Rows
		err  error
	)

	if rows, err = db.sqlDb.Query("SHOW TABLES"); err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return err
		}
		query := fmt.Sprintf("DESC %s", name)

		irows, err := db.sqlDb.Query(query)
		if err != nil {
			return err
		}

		columns, _ := irows.Columns()
		values := make([]sql.RawBytes, len(columns))
		scanArgs := make([]interface{}, len(values))
		for i := range values {
			scanArgs[i] = &values[i]
		}

		tableDesc := TableDesc{name: name}
		tableDesc.Fields = make([][]string, 0)

		for irows.Next() {
			irows.Scan(scanArgs...)
			idx_field := IndexSlice(columns, "Field")
			idx_type := IndexSlice(columns, "Type")

			field, fieldtype := string(*scanArgs[idx_field].(*sql.RawBytes)),
				string(*scanArgs[idx_type].(*sql.RawBytes))

			d := []string{field, fieldtype}
			tableDesc.Fields = append(tableDesc.Fields, d)
		}
		myfile.WriteStruct(tableDesc)
	}
	if err := rows.Err(); err != nil {
		log.Fatalln(err)
	}
	myfile.FormatFile()
	return nil
}

func IndexSlice(s []string, t string) int {
	for i, v := range s {
		if v == t {
			return i
		}
	}
	return -1
}
