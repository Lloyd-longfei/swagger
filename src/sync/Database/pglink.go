package Database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"sync/Helpers"
	"log"
)

var container interface{}


/**
	Query for data

 */
func Querys( opt string, sqls string) *[]interface{} {
	db, err := sql.Open("postgres", opt)
	Helpers.CheckError(err)
	err = db.Ping()
	Helpers.CheckError(err)
	rows, err := db.Query(sqls)
	Helpers.CheckError(err)
	columns, err := rows.Columns()
	colNum := len(columns)
	var values = make([]interface{}, colNum)
	for i, _ := range values {
		values[i] = &container
	}
	var arrLists []interface{}
	for rows.Next() {
		rows.Scan(values...)
		mapInstance := make(map[string]interface{})
		for i, colName := range columns {
			mapInstance[colName] = *(values[i].(*interface{}))
		}
		arrLists = append(arrLists, mapInstance)
	}
	defer db.Close()
	return &arrLists
}

/**
用于 insert update delete
 */
func Exec(opt string, sqls string) bool {
	db, err := sql.Open("postgres", opt)
	Helpers.CheckError(err)
	err = db.Ping()
	Helpers.CheckError(err)
	res, err := db.Exec(sqls)
	Helpers.CheckError(err)
	log.Print(res)
	defer db.Close()
	return true
}
