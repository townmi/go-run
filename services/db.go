package services

import (
	"database/sql"
	"reflect"
	"go-run/config"
	_ "go-run/model"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/go-sql-driver/mysql"
)

var DbLink string

type stockDBModel struct {
	StockId        string
	StockName      string
	StockChinaName string
}

func init() {

	// set db link
	switch config.Env.SQL.NAME {
	case "sqlite3":
		DbLink = config.Env.SQL.LOCAL
	case "mysql":
		DbLink = config.Env.SQL.USER + ":" + config.Env.SQL.PASSWORD + "@tcp(" + config.Env.SQL.HOST + ":" + config.Env.SQL.PORT + ")/" + config.Env.SQL.DATABASE + "?tls=skip-verify&autocommit=true"
	default:
		DbLink = config.Env.SQL.LOCAL
	}
}

func Select(query string, model interface{}, cond ...interface{}) []interface{} {

	db, err := sql.Open(config.Env.SQL.NAME, DbLink)
	config.CheckError(err, "open db fail\n")
	defer db.Close()

	stmt, err := db.Prepare(query)
	config.CheckError(err, "prepare sql for ---- ****** ----> " + query + " <---- ****** ----fail\n")

	rows, err := stmt.Query(cond...)
	config.CheckError(err, "Query db fail\n")

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

func Insert(query string, cond ...interface{}) interface{} {

	db, err := sql.Open(config.Env.SQL.NAME, DbLink)
	config.CheckError(err, "open db fail\n")
	defer db.Close()

	stmt, err := db.Prepare(query)
	config.CheckError(err, "prepare sql for ---- ****** ----> " + query + " <---- ****** ----fail\n")

	res, err := stmt.Exec(cond...)
	config.CheckError(err, "insert values fail\n")

	defer stmt.Close()

	lastId, err := res.LastInsertId()
	config.CheckError(err, "get last Id fail\n")

	rowCnt, err := res.RowsAffected()
	config.CheckError(err, "get rows fail")

	result := make([]interface{}, 0)
	result = append(result, lastId, rowCnt)

	return result
}

//事物处理
func InsertDataDBTx(insertSql string, argsList []interface{}) string {

	db, err := sql.Open(config.Env.SQL.NAME, DbLink)
	config.CheckError(err, "open db fail\n")
	defer db.Close()

	tx, err := db.Begin()
	config.CheckError(err, "Begin db fail\n")

	stmtInsert, err := tx.Prepare(insertSql)
	config.CheckError(err, "Begin db fail\n")

	for i := 0; i < len(argsList); i++ {
		_, err = stmtInsert.Exec(argsList[i])
		config.CheckError(err, "Exec db fail\n")
	}
	stmtInsert.Close()

	err = tx.Commit()
	config.CheckError(err, "Commit db fail\n")

	return "ok"
}