package route

import (
	"net/http"
	"database/sql"
	"log"
	"fmt"
	_ "go-run/config"
	_ "go-run/model"
	_ "github.com/go-sql-driver/mysql"
	"go-run/config"
)

type searchParams struct {
}

func GetSearch(w http.ResponseWriter, r *http.Request) {

	dbLink := config.Env.SQL.USER + ":" + config.Env.SQL.PASSWORD + "@tcp(" + config.Env.SQL.HOST + ":" + config.Env.SQL.PORT + ")/stock?tls=skip-verify&autocommit=true"

	db, err := sql.Open(config.Env.SQL.NAME, dbLink)

	if err != nil {
		log.Fatal(err)
	}
	stockid := 600007

	rows, err := db.Query("SELECT STOCKNAME, STOCKCHINANAME FROM stockCollections WHERE STOCKID=?", stockid)

	if err != nil {
		log.Fatal(err)
	}

	db.Close()

	defer rows.Close()

	for rows.Next() {
		var stockname string
		var stockchinaname string
		if err := rows.Scan(&stockname, &stockchinaname); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s & %v is %d\n", stockname, stockchinaname, stockid)
	}

	w.Write([]byte("{a:12}"))
}
