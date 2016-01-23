package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var Maindb *MySqlDbClient

type MySqlDbClient struct {
	config *ConfigParams
	sqlDb  *sql.DB
}

func initializeDB(configparams ConfigParams) {
	log.Printf("--> %v \n", configparams)
	Maindb = NewMySqlDbClient(&configparams)
}

func NewMySqlDbClient(config *ConfigParams) *MySqlDbClient {
	return &MySqlDbClient{config: config}
}

func (db *MySqlDbClient) Open() error {
	var err error
	connectionString := fmt.Sprintf("%s:%s@/%s", db.config.user, db.config.password, db.config.database)
	db.sqlDb, err = sql.Open("mysql", connectionString)
	if err != nil {
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
