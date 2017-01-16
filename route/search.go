package route

import (
	"net/http"
	_ "strconv"
	"time"
	_ "fmt"
	"go-run/config"
	DB "go-run/services"
	"encoding/json"
	_ "io/ioutil"
)

type search struct {
	VALUE string `json:"value"`
}

type searchId struct {
	StockId string
}

func GetSearch(w http.ResponseWriter, r *http.Request) {

	config.SetCORS(w)

	model := struct {
		ID             int
		StockName      string
		StockId        string
		STOCKCHINANAME string
		CREATEDAT      time.Time
		UPDATEAT       time.Time
	}{}

	sqlString := "SELECT ID, StockName, StockId, STOCKCHINANAME, CREATEDAT, UPDATEAT FROM stockCollections"

	data := DB.Select(sqlString, &model)

	send, _ := json.Marshal(data)

	w.Write([]byte(string(send)))

}

func PostSearch(w http.ResponseWriter, r *http.Request) {

	config.SetCORS(w)

	// t args from client
	var t search

	r.ParseForm()

	//decoder := json.NewDecoder(r.Body)
	//err := decoder.Decode(&t)
	//config.CheckError(err, "Json decode r body fail")

	t.VALUE = r.Form.Get("value")

	defer r.Body.Close()

	model := struct {
		StockId        string
		StockName      string
		StockChinaName string
	}{}

	sqlString := "SELECT STOCKID, STOCKNAME, STOCKCHINANAME FROM stockCollections WHERE STOCKID LIKE '%" + t.VALUE + "%'"

	data := DB.Select(sqlString, &model)

	send, _ := json.Marshal(data)

	w.Write([]byte(string(send)))

}
