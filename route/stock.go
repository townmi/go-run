package route

import (
	"net/http"
	"encoding/json"
	"go-run/config"
	"io/ioutil"
	DB "go-run/services"
	"github.com/robertkrimen/otto"
	_ "bytes"
	"fmt"
)

type stockList struct {
	VAL  string
	VAL3 string
	VAL2 string
}

var stockLists []stockList

func GetStock(w http.ResponseWriter, r *http.Request) {

	config.SetCORS(w)

	resp, err := http.Get("http://www.sse.com.cn/js/common/ssesuggestdataAll.js")
	config.CheckError(err, "http get stocks fail")

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	config.CheckError(err, "ioutil read data fail")

	vm := otto.New()
	vm.Run(string(body) + "\n;var stockList = JSON.stringify(get_alldata());")

	value, err := vm.Get("stockList")
	config.CheckError(err, "vm get javascript data fail")

	v, _ := value.ToString()
	errJson := json.Unmarshal([]byte(v), &stockLists)
	config.CheckError(errJson, "Json Unmarshal fail")

	model := struct {
		StockId string
	}{}

	sqlString := "SELECT StockId FROM stockLists"

	data := DB.Select(sqlString, &model)

	for _, v := range data {
		fmt.Println(v)
	}

	send, _ := json.Marshal(data)

	w.Write([]byte(string(send)))

}