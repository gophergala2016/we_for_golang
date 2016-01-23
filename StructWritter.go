package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"unicode"
)

type FileIO struct {
	fullfile    string
	fileHandler *os.File
}

type TableDesc struct {
	name   string
	Fields [][]string
}

var myfile FileIO

func NewFileIO(config ConfigParams) error {
	var err error
	myfile = FileIO{fullfile: config.fileforStruct}
	myfile.fileHandler, err = os.OpenFile(config.fileforStruct,
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
	l1 := fmt.Sprintf("type %s struct { \n", UpcaseInitial(tableDesc.name))

	f.fileHandler.WriteString(l1)
	for _, v := range tableDesc.Fields {
		_field := UpcaseInitial(v[0])
		_type := getGoTypes(v[1])
		fieldStr := fmt.Sprintf("%s %s \n", _field, _type)
		f.fileHandler.WriteString(fieldStr)
	}
	f.fileHandler.WriteString("}")
	f.fileHandler.WriteString("\n")
}

func UpcaseInitial(str string) string {
	if str == "" {
		return ""
	}
	strInRune := []rune(str)
	strInRune[0] = unicode.ToUpper(strInRune[0])
	return string(strInRune)
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
	} else if strings.Contains(t, "text") {
		return "string"
	}

	return t
}

func (f *FileIO) FormatFile() error {
	cmd := "go"
	args := []string{"fmt", f.fullfile}

	if err := exec.Command(cmd, args...).Run(); err != nil {
		return err
	}
	return nil
}
