package route

import (
	"net/http"
	"database/sql"
	"log"
	_ "go-run/model"
	_ "github.com/go-sql-driver/mysql"
	"go-run/config"
)

type searchParams struct {
}

func GetSearch(w http.ResponseWriter, r *http.Request) {

	dbLink := config.Env.SQL.USER + ":" + config.Env.SQL.PASSWORD + "@tcp(" + config.Env.SQL.HOST + ":" + config.Env.SQL.PORT + ")/" + config.Env.SQL.DATABASE + "?tls=skip-verify&autocommit=true"

	db, err := sql.Open(config.Env.SQL.NAME, dbLink)
	config.CheckError(err)

	id := "21"

	rows, err := db.Query("SELECT title, body FROM art WHERE id=?", id)
	config.CheckError(err)

	db.Close()

	defer rows.Close()

	var title string
	var body string

	for rows.Next() {

		if err := rows.Scan(&title, &body); err != nil {
			log.Fatal(err)
		}
	}

	w.Write([]byte("{title:" + title + ", body: " + body + ", id:" + id + "}"))
}
