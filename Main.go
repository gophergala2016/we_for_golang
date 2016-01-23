package main

import (
	"flag"
	"log"
	"os"
)

type ConfigParams struct {
	user     string
	password string
	database string
}

var configparams ConfigParams

func init() {
	configparams = ConfigParams{}
}

func main() {

	flag.StringVar(&configparams.user, "u", "", "User")
	flag.StringVar(&configparams.password, "p", "", "Password")
	flag.StringVar(&configparams.database, "d", "", "Name of Database")
	flag.Parse()

	if (configparams.user == "") || (configparams.database == "") || (configparams.password == "") {
		log.Println(configparams)
		log.Println("In sufficient paramas")
		os.Exit(0)
	}
	initializeDB(configparams)
	NewFileIO()

	err := db.Open()
	if err != nil {
		log.Println("-------------------------------")
		log.Fatalln(err.Error())
	}
	defer db.Close()

	//	db.ShowTables()
	db.GenerateSchemaFile()
	//initStructWritter()

}
