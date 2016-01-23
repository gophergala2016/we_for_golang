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
	log.Printf("--> %v \n", configparams)
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

	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", name)
	}
	if err := rows.Err(); err != nil {
		fmt.Println(err)
	}
}
