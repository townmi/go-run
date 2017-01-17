package route

import (
	"net/http"
	"encoding/json"
	"go-run/config"
	"io/ioutil"
	DB "go-run/services"
	"github.com/robertkrimen/otto"
	"reflect"
	"fmt"
)

type stockList struct {
	STOCKID        string
	STOCKNAME      string
	STOCKCHINANAME string
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

	result, _ := vm.Run(string(body) + "\n;console.log(JSON.stringify(get_alldata()));")
	fmt.Println(result)

	s := reflect.ValueOf(&otto.Value{}).Elem()
	len := s.NumField()

	fmt.Println( s.FieldByName("value"), len)

	//errJson := json.Unmarshal([]byte("{}"), &stockLists)
	//config.CheckError(errJson, "Json Unmarshal fail")

	model := struct {
		ID int
	}{}

	sqlString := "SELECT count(*) AS `count` FROM `stockLists` AS `stockList` WHERE `stockList`.`STOCKID` = '600012' AND `stockList`.`STOCKNAME` = 'msyx'"

	data := DB.Select(sqlString, &model)

	send, _ := json.Marshal(data)

	w.Write([]byte(string(send)))

}