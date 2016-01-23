package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

type FileIO struct {
	fullfile    string
	fileHandler *os.File
}

type TableDesc struct {
	name string
	//	Field map[string]string
	Fields [][]string
}

var myfile FileIO

func NewFileIO() error {
	var err error
	fileName := "./test.go"
	myfile = FileIO{fullfile: fileName}
	myfile.fileHandler, err = os.OpenFile(fileName,
		os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)

	if err != nil {
		return err
	}
	myfile.WriteHeaders()
	return nil
}
func (f *FileIO) WriteHeaders() {
	f.fileHandler.WriteString("package main \n")
	f.fileHandler.WriteString(`import("fmt")`)
	f.fileHandler.WriteString("\n")
}

func (f *FileIO) WriteStruct(tableDesc TableDesc) {
	f.fileHandler.WriteString("\n")
	l1 := fmt.Sprintf("type %s struct { \n", tableDesc.name)

	f.fileHandler.WriteString(l1)
	for _, v := range tableDesc.Fields {

		matched_Type := getGoTypes(v[1])
		fieldStr := fmt.Sprintf("%s %s \n", v[0], matched_Type)
		f.fileHandler.WriteString(fieldStr)
	}
	f.fileHandler.WriteString("}")
	f.fileHandler.WriteString("\n")
}

func getGoTypes(t string) string {

	if strings.Contains(t, "tinyint") {
		return "bool"
	} else if strings.Contains(t, "int") {
		return "int"
	} else if strings.Contains(t, "varchar") {
		return "string"
	} else if strings.Contains(t, "datetime") {
		return "time.Time"
	} else if strings.Contains(t, "decimal") {
		return "float64"
	}
	return t
}

func (f *FileIO) FormatFile() {
	cmd := "go"
	args := []string{"fmt", f.fullfile}

	if err := exec.Command(cmd, args...).Run(); err != nil {
		fmt.Println("ERRRRRRRRRRRRRRRRRRRRRRRR")
		log.Fatal(err)
	}
}

//func initStructWritter() {
//	var err error
//	//	fileName := "/home/synerzip/workspace/go/src/VimleshS/gophergala/we_for_golang/MySqlStructs.go"
//	fileName := "./test.go"

//	fileHandler, err = os.OpenFile(fileName,
//		os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
//	if err != nil {
//		return
//	}
//	defer fileHandler.Close()
//	fmt.Println("created..")
//}
