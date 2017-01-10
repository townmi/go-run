package route

import (
	"net/http"
	"database/sql"
	"log"
	_ "go-run/model"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/go-sql-driver/mysql"
	"go-run/config"
	_ "fmt"
)

type searchParams struct {
}

func GetSearch(w http.ResponseWriter, r *http.Request) {

	var dbLink string

	switch config.Env.SQL.NAME {
	case "sqlite3":
		dbLink = config.Env.SQL.LOCAL
	case "mysql":
		dbLink = config.Env.SQL.USER + ":" + config.Env.SQL.PASSWORD + "@tcp(" + config.Env.SQL.HOST + ":" + config.Env.SQL.PORT + ")/" + config.Env.SQL.DATABASE + "?tls=skip-verify&autocommit=true"
	default:
		dbLink = config.Env.SQL.LOCAL
	}

	db, err := sql.Open(config.Env.SQL.NAME, dbLink)
	config.CheckError(err)

	uid := "1"

	rows, err := db.Query("SELECT * FROM userinfo")
	config.CheckError(err)

	db.Close()

	defer rows.Close()

	var username string
	var department string
	var created string

	for rows.Next() {

		if err := rows.Scan(&uid, &username, &department, &created); err != nil {
			log.Fatal(err)
		}
	}

	w.Write([]byte("{title:" + username + ", body: " + department + ", id:" + uid + "}"))
}
