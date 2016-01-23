package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
)

type ConfigParams struct {
	user          string
	password      string
	database      string
	fileforStruct string
}

var configparams ConfigParams

func init() {
	configparams = ConfigParams{}
}

func main() {

	flag.StringVar(&configparams.user, "u", "", "User")
	flag.StringVar(&configparams.password, "p", "", "Password")
	flag.StringVar(&configparams.database, "d", "", "Name of Database")
	flag.StringVar(&configparams.fileforStruct, "f", "./Test.go", "Specify file for Golang Struct")
	flag.Parse()
	
	ext := filepath.Ext(configparams.fileforStruct)
	if ext == "" || ext != ".go" {
		configparams.fileforStruct = configparams.fileforStruct +  ".go"
	}

	if (configparams.user == "") || (configparams.database == "") || (configparams.password == "") {
		log.Println(configparams)
		log.Println("In sufficient parameter")
		os.Exit(0)
	}
	initializeDB(configparams)
	NewFileIO(configparams)

	err := db.Open()
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer db.Close()
	if err := db.GenerateSchemaFile(); err != nil {
		log.Fatal(err)
	}
	log.Printf("Structure written to %s \n", configparams.fileforStruct)
}
