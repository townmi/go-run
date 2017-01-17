package services

import (
	"database/sql"
	"reflect"
	"go-run/config"
	_ "go-run/model"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/go-sql-driver/mysql"
)

var dbLink string

func init() {

	// set db link
	switch config.Env.SQL.NAME {
	case "sqlite3":
		dbLink = config.Env.SQL.LOCAL
	case "mysql":
		dbLink = config.Env.SQL.USER + ":" + config.Env.SQL.PASSWORD + "@tcp(" + config.Env.SQL.HOST + ":" + config.Env.SQL.PORT + ")/" + config.Env.SQL.DATABASE + "?tls=skip-verify&autocommit=true"
	default:
		dbLink = config.Env.SQL.LOCAL
	}
}

func Select(query string, model interface{}, cond ...interface{}) []interface{} {

	db, err := sql.Open(config.Env.SQL.NAME, dbLink)
	config.CheckError(err, "open db fail")
	defer db.Close()

	stmt, err := db.Prepare(query)
	config.CheckError(err, "prepare sql for ---- ****** ----> " + query + " <---- ****** ----fail")
	defer stmt.Close()

	rows, err := stmt.Query(cond...)
	config.CheckError(err, "Query db fail")
	defer rows.Close()

	result := make([]interface{}, 0)

	s := reflect.ValueOf(model).Elem()
	len := s.NumField()
	rowCells := make([]interface{}, len)
	for i := 0; i < len; i++ {
		rowCells[i] = s.Field(i).Addr().Interface()
	}

	for rows.Next() {
		err = rows.Scan(rowCells...)
		if err != nil {
			panic(err)
		}
		result = append(result, s.Interface())
	}
	return result
}

//func Insert(query string, cond ...interface{}) interface{} {
//
//	db, err := sql.Open(config.Env.SQL.NAME, dbLink)
//	config.CheckError(err, "open db fail")
//	defer db.Close()
//
//	stmt, err := db.Prepare(query)
//	config.CheckError(err, "prepare sql for ---- ****** ----> " + query + " <---- ****** ----fail")
//
//	res, err := stmt.Exec(cond...)
//	config.CheckError(err, "insert values fail")
//
//	//lastId, err := res.LastInsertId()
//	//config.CheckError(err, "get last Id fail")
//
//	rowCnt, err := res.RowsAffected()
//	config.CheckError(err, "get rows fail")
//
//	return  &rowCnt
//}