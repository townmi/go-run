package route

import (
	"net/http"
	"database/sql"
	_ "go-run/model"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/go-sql-driver/mysql"
	"go-run/config"
	_ "fmt"
	"time"
	"encoding/json"
	_ "io/ioutil"
)

type stock struct {
	ID             int64
	STOCKID        string
	STOCKNAME      string
	STOCKCHINANAME string
	CREATEDAT      time.Time
	UPDATEAT       time.Time
}

type search struct {
	VALUE string
}

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

func GetSearch(w http.ResponseWriter, r *http.Request) {

	db, err := sql.Open(config.Env.SQL.NAME, dbLink)
	config.CheckError(err)

	rows, err := db.Query("SELECT * FROM stockCollections")
	config.CheckError(err)

	db.Close()

	defer rows.Close()

	var data [] stock

	for rows.Next() {
		var id int64
		var stockId string
		var stockName string
		var stockChinaName string
		var created time.Time
		var updated time.Time

		err := rows.Scan(&id, &stockId, &stockName, &stockChinaName, &created, &updated)
		config.CheckError(err)

		val := stock{
			ID:             id,
			STOCKID:        stockId,
			STOCKNAME:      stockName,
			STOCKCHINANAME: stockChinaName,
			CREATEDAT:      created,
			UPDATEAT:       updated,
		}

		data = append(data, val)
	}
	send, _ := json.Marshal(data)
	w.Write([]byte(string(send)))

}

func PostSearch(w http.ResponseWriter, r *http.Request) {
	// t args from client
	var t search

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&t)
	config.CheckError(err)

	defer r.Body.Close()

	db, err := sql.Open(config.Env.SQL.NAME, dbLink)
	config.CheckError(err)

	rows, err := db.Query("SELECT * FROM stockCollections")
	config.CheckError(err)

	db.Close()

	defer rows.Close()

	var data [] stock

	for rows.Next() {
		var id int64
		var stockId string
		var stockName string
		var stockChinaName string
		var created time.Time
		var updated time.Time

		err := rows.Scan(&id, &stockId, &stockName, &stockChinaName, &created, &updated)
		config.CheckError(err)

		val := stock{
			ID:             id,
			STOCKID:        stockId,
			STOCKNAME:      stockName,
			STOCKCHINANAME: stockChinaName,
			CREATEDAT:      created,
			UPDATEAT:       updated,
		}

		data = append(data, val)
	}
	send, _ := json.Marshal(data)
	w.Write([]byte(string(send)))

}
