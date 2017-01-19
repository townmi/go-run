package route

import (
	"net/http"
	"encoding/json"
	"go-run/config"
	"io/ioutil"
	DB "go-run/services"
	"github.com/robertkrimen/otto"
	_ "crypto/sha1"
	_ "io"
	"fmt"
	_ "reflect"
)

type stockList struct {
	VAL  string
	VAL3 string
	VAL2 string
}

type stockDBModel struct {
	StockId        string
	StockName      string
	StockChinaName string
}

var stockLists []stockList

func init() {

	//var list [2]stockList
	//var list2 [2]stockList
	//
	//hash := make(map[string]int)
	//
	//list[0] = stockList{VAL:"600000", VAL3:"pfyx", VAL2:"浦发银行"}
	//list[1] = stockList{VAL:"600018", VAL3:"sgjt", VAL2:"上港集团"}
	//
	//list2[0] = stockList{VAL:"600019", VAL3:"bggf", VAL2:"宝钢股份"}
	//list2[1] = stockList{VAL:"600000", VAL3:"pfyx", VAL2:"浦发银行"}
	//
	//for _, v := range list {
	//
	//	h := sha1.New()
	//
	//	s := v.VAL + v.VAL2
	//
	//	io.WriteString(h, s)
	//
	//	bs := h.Sum(nil)
	//	str := ByteToHex(bs)
	//
	//	hash[str] = 1
	//
	//}
	//
	//for _, v := range list2 {
	//
	//	h := sha1.New()
	//
	//	s := v.VAL + v.VAL2
	//
	//	io.WriteString(h, s)
	//
	//	bs := h.Sum(nil)
	//	str := ByteToHex(bs)
	//
	//	fmt.Println(hash[str])
	//
	//}
	//
	//fmt.Println(hash)
}

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

	sqlString := "SELECT STOCKID, STOCKNAME, STOCKCHINANAME FROM stockLists"
	data := DB.Select(sqlString, &stockDBModel{})

	hashMap := make(map[string]int)
	hashResult := make([]stockList, 0)

	for _, v := range data {

		sv := v.(stockDBModel)
		s := sv.StockId + sv.StockChinaName

		hashMap[s] = 1
	}

	for _, v := range stockLists {

		s := v.VAL + v.VAL2
		sv := hashMap[s]

		if sv == 0 {
			hashResult = append(hashResult, v)
		}
	}

	fmt.Println(hashResult)
	insertSlice := make([]stockDBModel, 0)
	for _, v := range hashResult {
		insertSlice = append(insertSlice, stockDBModel{v.VAL, v.VAL3, v.VAL2})
	}

	inserString := "INSERT INTO stockLists(STOCKID, STOCKNAME, STOCKCHINANAME) values(?,?,?)"

	insertData := DB.Insert(inserString, &insertSlice)


	send, _ := json.Marshal(insertData)

	w.Write([]byte(string(send)))

}