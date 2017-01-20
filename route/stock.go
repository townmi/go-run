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
	_ "reflect"
	"bytes"
	"fmt"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/go-sql-driver/mysql"
)

type stockList struct {
	VAL  string
	VAL3 string
	VAL2 string
}

type stockDBModel struct {
	STOCKID string
	OPENATCASH float64
	MIDCLOSEATCASH float64
	MIDOPENATCASH float64
	CLOSEATCASH float64
	TRADECOUNT float64
	STOCKCATEGORY string
	DATE string
}

type stockListDBModel struct {
	StockId        string
	StockName      string
	StockChinaName string
}

var stock [][]string

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

func GetStockList(w http.ResponseWriter, r *http.Request)  {
	config.SetCORS(w)

	sqlString := "SELECT STOCKID, STOCKNAME, STOCKCHINANAME FROM stockLists"
	data := DB.Select(sqlString, &stockListDBModel{})

	send, _ := json.Marshal(data)

	w.Write([]byte(string(send)))
}

func GetStock(w http.ResponseWriter, r *http.Request)  {

	config.SetCORS(w)

	sqlString := "SELECT STOCKID, OPENATCASH, MIDCLOSEATCASH, MIDOPENATCASH, CLOSEATCASH, TRADECOUNT, STOCKCATEGORY, DATE FROM stockCollections WHERE "
	data := DB.Select(sqlString, &stockDBModel{})

	send, _ := json.Marshal(data)

	w.Write([]byte(string(send)))
}



func ReFreshStockList(w http.ResponseWriter, r *http.Request) {

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
	data := DB.Select(sqlString, &stockListDBModel{})

	hashMap := make(map[string]int)
	hashResult := make([]stockList, 0)

	for _, v := range data {

		sv := v.(stockListDBModel)
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

	if len(hashResult) == 0 {
		w.Write([]byte(`{action: "nodiff"}`))
	} else {
		b := bytes.Buffer{}
		b.WriteString("INSERT INTO stockLists(STOCKID, STOCKNAME, STOCKCHINANAME) values")

		insertSlice := make([]interface{}, 0)
		for i, v := range hashResult {
			insertSlice = append(insertSlice, v.VAL, v.VAL3, v.VAL2)
			if i == len(hashResult) - 1 {
				b.WriteString("(?,?,?)")
			} else {
				b.WriteString("(?,?,?), ")
			}
		}

		rowCnt := DB.Insert(b.String(), insertSlice...)

		fmt.Println(insertSlice[0])

		send, _ := json.Marshal(rowCnt)

		w.Write([]byte(string(send)))
	}
}



func ReFreshStock(w http.ResponseWriter, r *http.Request) {

	config.SetCORS(w)
	//insertSlice := make([]interface{}, 0)

	type model struct {
		StockId string
	}

	sqlString := "SELECT STOCKID FROM stockLists"
	data := DB.Select(sqlString, &model{})

	b := bytes.Buffer{}
	b.WriteString("INSERT INTO stockCollections(DATE, OPENATCASH, MIDCLOSEATCASH, MIDOPENATCASH, CLOSEATCASH, TRADECOUNT, STOCKCATEGORY, STOCKID) values (?,?,?,?,?,?,?,?)")



	/**
	 * 事物
	 */

	db, err := sql.Open(config.Env.SQL.NAME, DB.DbLink)
	config.CheckError(err, "open db fail\n")
	defer db.Close()

	tx, err := db.Begin()
	config.CheckError(err, "Begin db fail\n")

	stmtInsert, err := tx.Prepare(b.String())
	config.CheckError(err, "Begin db fail\n")


	for _, v := range data {

		//if i < 3 {
			sv := v.(model)
			st := "http://web.ifzq.gtimg.cn/appstock/app/fqkline/get?_var=kline_dayqfq&param=sh" + sv.StockId + ",day,2017-01-01,2017-12-31,320,qfq&r="

			resp, err := http.Get(st)
			config.CheckError(err, "http get stocks fail\n")

			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			config.CheckError(err, "ioutil read data fail\n")

			vm := otto.New()
			vm.Run(string(body) + "\n;var yearData = JSON.stringify(kline_dayqfq.data.sh" + sv.StockId + ".qfqday);")

			value, err := vm.Get("yearData")
			config.CheckError(err, "vm get javascript data fail\n")

			v, _ := value.ToString()
			if v != "undefined" {
				errJson := json.Unmarshal([]byte(v), &stock)
				config.CheckError(errJson, "")

				for _, vj := range stock {
					if len(vj) == 6 {

						_, err = stmtInsert.Exec(vj[0], vj[1], vj[2], vj[3], vj[4], vj[5], "sh", sv.StockId)
						config.CheckError(err, "Exec db fail\n")

						//insertSlice = append(insertSlice, []interface{}{})
						//b.WriteString("(?,?,?,?,?,?,?,?), ")
					}

				}

			}

		//}
		//if i == len(data) - 1 {
		//	b.WriteString("(?,?,?,?,?,?,?,?)")
		//} else {
		//	b.WriteString("(?,?,?,?,?,?,?,?), ")
		//}

	}
	//
	//for _, v := range insertSlice {
	//
	//}

	//sb := b.String()
	//
	//sb = sb[0:len(sb) - 2]
	//
	//if len(requsetList) == 0 {
	//	w.Write([]byte(`{action: "nodiff"}`))
	//} else {
	//	b := bytes.Buffer{}
	//	b.WriteString("INSERT INTO stockLists(STOCKID, STOCKNAME, STOCKCHINANAME) values")
	//
	//	insertSlice := make([]interface{}, 0)
	//	for i, v := range hashResult {
	//		insertSlice = append(insertSlice, v.VAL, v.VAL3, v.VAL2)
	//		if i == len(hashResult) - 1 {
	//			b.WriteString("(?,?,?)")
	//		} else {
	//			b.WriteString("(?,?,?), ")
	//		}
	//	}
	//
	//	rowCnt := DB.Insert(b.String(), insertSlice...)
	//
	//	send, _ := json.Marshal(rowCnt)
	//
	//	w.Write([]byte(string(send)))
	//}
	//fmt.Println(insertSlice[0])
	//rowCnt := DB.Insert(b.String(), insertSlice...)

	//rowCnt := DB.InsertDataDBTx(b.String(), insertSlice)

	stmtInsert.Close()

	err = tx.Commit()

	//send, _ := json.Marshal(rowCnt)

	//w.Write([]byte(string(send)))

	w.Write([]byte("ok"))
}