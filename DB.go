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
	defer rows.Close()

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

func (db *MySqlDbClient) GenerateSchemaFile() error {

	//	rows, err := db.sqlDb.Query("SHOW TABLES")
	//	if err != nil {
	//		panic(err)
	//	}
	//	defer rows.Close()

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
		//fmt.Println(query)

		irows, err := db.sqlDb.Query(query)
		if err != nil {
			return err
		}

		columns, _ := irows.Columns()
		values := make([]sql.RawBytes, len(columns))
		//		fmt.Println(values)
		//		//scanArgs := make([]interface{}, len(values))
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

			//			fmt.Printf("%d : %15s  %15s |", IndexSlice(columns, "Type"),
			//				string(*scanArgs[idx_field].(*sql.RawBytes)),
			//				string(*scanArgs[idx_type].(*sql.RawBytes)))

			field, fieldtype := string(*scanArgs[idx_field].(*sql.RawBytes)),
				string(*scanArgs[idx_type].(*sql.RawBytes))

			d := []string{field, fieldtype}
			tableDesc.Fields = append(tableDesc.Fields, d)

		}
		myfile.WriteStruct(tableDesc)
	}
	if err := rows.Err(); err != nil {
		fmt.Println(err)
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

func GenerateStruct(columns []string, readData []interface{}) {

	//	fmt.Println("-------------------------")
	//	fmt.Println(IndexSlice(columns, "Typed"))
	//	fmt.Println("-------------------------")

	//	for i := 0; i < len(readData); i++ {
	//		idx_Field := IndexSlice(columns, "Field")
	//		idx_Type := IndexSlice(columns, "Type")

	//		row_Values := readData[i]

	//		fmt.Printf("%5s : %15s |", IndexSlice(columns, "Type"), string(*value.(*sql.RawBytes)))

	//	}

	fmt.Println("++++++++++++++++++++++++++++++++++")
	for idx, value := range readData {
		fmt.Printf("%5s : %15s |", columns[idx], string(*value.(*sql.RawBytes)))
		fmt.Println(idx)
		ttt := value.(*sql.RawBytes)
		fmt.Println(string(*ttt))
	}

	fmt.Println()

}

/*
func (db *MySqlDbClient) GetRowsAndDetail(rows *sql.Rows) (columns []string, rowdata [](map[string]string), err error) {

	if columns, err = rows.Columns(); err != nil {
		return nil, nil, err
	}

	columndata := make([]interface{}, len(columns))
	for i := range columndata {
		columndata[i] = new(interface{})
	}

	rowdata = make([]map[string]string, 0)

	for rows.Next() {
		rows.Scan(columndata...)

		data := make(map[string]string)

		for i, value := range columndata {

		}
		//		for i, val := range columns {
		//			fmt.Println(i)

		//		}
		//		fmt.Println()

	}

	//	for rows.Next() {
	//		var name string
	//		if err := rows.Scan(&name); err != nil {
	//			return nil, nil, err
	//		}
	//		fmt.Printf("%s\n", name)
	//	}

	if err := rows.Err(); err != nil {
		return nil, nil, err
	}

	return columns, rowdata, nil
}

*/
