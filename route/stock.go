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
	"bytes"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"strconv"
)

type httpStockList struct {
	scriptName     string
	scriptFunc     string
	scriptCon      string
	scriptConShort string
	scriptOrg      string
}
type stockList struct {
	VAL  string
	VAL3 string
	VAL2 string
}

type stockDBModel struct {
	CloseAtCash float64
	Date        int64
}

type stockListDBModel struct {
	StockId        string
	StockName      string
	StockChinaName string
}

type stockInfo struct {
	stockId       string
	stockConShort string
	stockOrg      string
	startDate     int64
	endDate       int64
}

func (R *stockInfo) reSetStockId(v string) {
	R.stockId = v
}
func (R *stockInfo) reSetStockConShort(v string) {
	R.stockConShort = v
}
func (R *stockInfo) reSetStockOrg(v string) {
	R.stockOrg = v
}
func (R *stockInfo) reSetStartDate(v string) {
	R.startDate = config.GetUnix(v)
}
func (R *stockInfo) reSetEndDate(v string) {
	R.endDate = config.GetUnix(v)
}

var stock [][]string

var stockLists []stockList

func GetStockList(w http.ResponseWriter, r *http.Request) {

	config.SetCORS(w)

	sqlString := "SELECT STOCKID, STOCKNAME, STOCKCHINANAME FROM stockLists"
	data := DB.Select(sqlString, &stockListDBModel{})

	send, _ := json.Marshal(data)

	w.Write([]byte(string(send)))
}

func GetStock(w http.ResponseWriter, r *http.Request) {

	config.SetCORS(w)

	selectObj := stockInfo{
		stockId: "000001",
		stockConShort: "Stock",
		stockOrg: "sh",
		startDate: config.GetUnix("2016-01-01"),
		endDate: config.GetUnix("2017-12-31"),
	}

	r.ParseForm()

	stockId := r.Form.Get("stockId")
	if stockId != "" {
		selectObj.reSetStockId(stockId)
	}
	stockConShort := r.Form.Get("stockConShort")
	if stockConShort != "" {
		selectObj.reSetStockConShort(stockConShort)
	}
	stockOrg := r.Form.Get("stockOrg")
	if stockOrg != "" {
		selectObj.reSetStockOrg(stockOrg)
	}
	startDate := r.Form.Get("startDate")
	if startDate != "" {
		selectObj.reSetStartDate(startDate)
	}
	endDate := r.Form.Get("endDate")
	if endDate != "" {
		selectObj.reSetEndDate(endDate)
	}

	sD := strconv.FormatInt(selectObj.startDate, 10)
	eD := strconv.FormatInt(selectObj.endDate, 10)
	sqlString := "SELECT CLOSEATCASH, DATE FROM `stockCollections` s WHERE s.STOCKUNIQUE = (SELECT STOCKUNIQUE FROM `stockLists` sl WHERE sl.STOCKID = '" + selectObj.stockId + "' AND sl.STOCKORG = '" + selectObj.stockOrg + "' AND sl.STOCKCONSHORT = '" + selectObj.stockConShort + "') AND s.DATE BETWEEN '" + sD + "' AND '" + eD + "'"

	//sqlString := "SELECT STOCKID, OPENATCASH, CLOSEATCASH, MAXATCASH, MINATCASH, TRADECOUNT, DATE FROM `stockCollections` s WHERE s.STOCKID = '"+selectObj.stockId+"'"

	data := DB.Select(sqlString, &stockDBModel{})

	send, _ := json.Marshal(data)

	w.Write([]byte(string(send)))
}

func ReFreshStockList(w http.ResponseWriter, r *http.Request) {

	config.SetCORS(w)

	lists := make([]httpStockList, 0)

	lists = append(lists, httpStockList{scriptName:"ssesuggestdata", scriptFunc:"get_data", scriptCon:"股票", scriptConShort:"Stock", scriptOrg:"sh"})
	lists = append(lists, httpStockList{scriptName:"ssesuggestfunddata", scriptFunc:"get_funddata", scriptCon:"基金", scriptConShort:"Fund", scriptOrg:"sh"})
	lists = append(lists, httpStockList{scriptName:"ssesuggestEbonddata", scriptFunc:"get_ebonddata", scriptCon:"可转换债券", scriptConShort:"Ebond", scriptOrg:"sh"})
	lists = append(lists, httpStockList{scriptName:"ssesuggestTbonddata", scriptFunc:"get_tbonddata", scriptCon:"国债/贴债", scriptConShort:"Tbond", scriptOrg:"sh"})

	result := runBackEndGetStockList(lists)

	send, _ := json.Marshal(result)

	w.Write([]byte(string(send)))

}

func ReFreshStock(w http.ResponseWriter, r *http.Request) {

	config.SetCORS(w)

	//r.ParseForm()
	//
	////decoder := json.NewDecoder(r.Body)
	////err := decoder.Decode(&t)
	////config.CheckError(err, "Json decode r body fail")
	//
	//t := r.Form.Get("s")
	//
	//fmt.Println(r.Form)
	//
	//w.Write([]byte(t))
	//
	//return

	count := 0
	year := [...]string{"1990", "1991", "1992", "1993", "1994", "1995", "1996", "1997", "1998", "1999", "2000", "2001", "2002", "2003", "2004", "2005", "2006", "2007", "2008", "2009", "2010", "2011", "2012", "2013", "2014", "2015", "2016", "2017"}

	type model struct {
		StockId     string
		StockUnique string
	}

	sqlString := "SELECT STOCKID, STOCKUNIQUE FROM `stockLists` WHERE `stockLists`.STOCKCONSHORT = 'Stock'"
	//sqlString := "SELECT STOCKID, STOCKUNIQUE FROM `stockLists` WHERE `stockLists`.STOCKCONSHORT = 'Stock' AND `stockLists`.STOCKID = '000001'"
	data := DB.Select(sqlString, &model{})

	b := bytes.Buffer{}
	b.WriteString("INSERT INTO stockCollections(STOCKID, OPENATCASH, CLOSEATCASH, MAXATCASH, MINATCASH, TRADECOUNT, DATE, STOCKUNIQUE, STOCKCOLLECTIONUNIQUE) values (?,?,?,?,?,?,?,?,?)")

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

	for _, y := range year {
		for _, v := range data {

			sv := v.(model)
			st := "http://web.ifzq.gtimg.cn/appstock/app/fqkline/get?_var=kline_dayqfq&param=sh" + sv.StockId + ",day," + y + "-01-01," + y + "-12-31,320,qfq&r="

			resp, err := http.Get(st)
			config.CheckError(err, "http get stocks fail\n")

			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			config.CheckError(err, "ioutil read data fail\n")

			vm := otto.New()
			var runScript string
			if sv.StockId == "000001" {
				runScript = string(body) + "\n;var yearData = JSON.stringify(kline_dayqfq.data.sh" + sv.StockId + ".day);"
			} else {
				runScript = string(body) + "\n;var yearData = JSON.stringify(kline_dayqfq.data.sh" + sv.StockId + ".qfqday);"
			}
			vm.Run(runScript)

			value, err := vm.Get("yearData")
			config.CheckError(err, "vm get javascript data fail\n")

			v, _ := value.ToString()

			if v != "undefined" {
				errJson := json.Unmarshal([]byte(v), &stock)
				config.CheckError(errJson, "")
				for _, vj := range stock {
					if len(vj) == 6 {
						timeUnix := config.GetUnix(vj[0])
						scu := sv.StockUnique + vj[0]
						_, err = stmtInsert.Exec(sv.StockId, vj[1], vj[2], vj[3], vj[4], vj[5], timeUnix, sv.StockUnique, scu)
						if err == nil {
							count ++
						} else {
							fmt.Println(err)
						}
					}

				}

			}

		}
	}

	stmtInsert.Close()

	err = tx.Commit()

	result := struct {
		status string
		count  string
	}{"success", string(count)}

	send, _ := json.Marshal(result)

	w.Write([]byte(string(send)))

	return

	//w.Write([]byte("ok"))
}

func CheckStockList(w http.ResponseWriter, r *http.Request) {
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

	send, _ := json.Marshal(hashResult)

	w.Write([]byte(string(send)))

	return

}

func runBackEndGetStockList(scripts []httpStockList) interface{} {

	result := make([]interface{}, 0)
	for _, v := range scripts {

		resp, err := http.Get("http://www.sse.com.cn/js/common/" + v.scriptName + ".js")
		config.CheckError(err, "http get stocks fail")

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		config.CheckError(err, "ioutil read data fail")

		vm := otto.New()
		vm.Run(string(body) + "\n;var stockList = JSON.stringify(" + v.scriptFunc + "());")

		value, err := vm.Get("stockList")
		config.CheckError(err, "vm get javascript data fail")

		val, _ := value.ToString()
		errJson := json.Unmarshal([]byte(val), &stockLists)
		config.CheckError(errJson, "Json Unmarshal fail")

		if len(stockLists) == 0 {
			result = append(result, nil)
		} else {

			b := bytes.Buffer{}
			b.WriteString("INSERT ignore INTO stockLists(STOCKID, STOCKCHINANAME, STOCKNAME, STOCKUNIQUE, STOCKCON, STOCKCONSHORT, STOCKORG) values")

			insertSlice := make([]interface{}, 0)
			for i, vs := range stockLists {
				unique := v.scriptOrg + v.scriptConShort + vs.VAL
				insertSlice = append(insertSlice, vs.VAL, vs.VAL2, vs.VAL3, unique, v.scriptCon, v.scriptConShort, v.scriptOrg)
				if i == len(stockLists) - 1 {
					b.WriteString("(?,?,?,?,?,?,?)")
				} else {
					b.WriteString("(?,?,?,?,?,?,?), ")
				}
			}

			rowCnt := DB.Insert(b.String(), insertSlice...)

			result = append(result, rowCnt)
		}

	}
	return result
}

/**
 ***********************************************************************
 * current stock data
 *
 *
 */

func ReFreshCurrentStock(w http.ResponseWriter, r *http.Request)  {
	config.SetCORS(w)
}