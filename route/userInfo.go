package route

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	"github.com/go-ini/ini"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
)

// UserDataBase is conf
type UserDataBase struct {
	USERNAME string
	PASSWORD string
	HOSTNAME string
	DATABASE string
	PORT     string
}

// User is struct model
type User struct {
	id              string
	displayName     string
	password        string
	role            string
	email           string
	mobile          string
	locationLat     string
	locationLon     string
	subscribe       int
	cityId          string
	secret          string
	miniWechatId    string
	WechatId        string
	level           int
	createdAt       time.Time
	bindMobileAt    time.Time
	updatedAt       time.Time
	consignee       string
	shippingAddress string
	receivingPhone  string
	isPolicy        int
	Postcode        string
	unionid         int
}

// Result is
type Result struct {
	displayName string
	id          string
}

var (
	db     UserDataBase
	dbLink string
)

// GetUserInfo is router for GetUserInfo
func GetUserInfo(c echo.Context) error {

	id := c.Param("id")

	b := bytes.Buffer{}
	b.WriteString(db.USERNAME)
	b.WriteString(":")
	b.WriteString(db.PASSWORD)
	b.WriteString("@tcp(")
	b.WriteString(db.HOSTNAME)
	b.WriteString(":")
	b.WriteString(db.PORT)
	b.WriteString(")/")
	b.WriteString(db.DATABASE)
	b.WriteString("?charset=utf8&parseTime=True&loc=Local")
	dbLink = b.String()

	db, err := gorm.Open("mysql", dbLink)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	db.DB().Ping()
	// db.AutoMigrate(&User{})
	db.LogMode(true)
	// var users = make([]User, 3)
	// rows, err := db.Debug().Table("user").Where("id = ?", id).Rows()
	rows, err := db.Table("user").Where("id = ?", id).Rows() // (*sql.Row)
	// rows, err := db.Table("user").Limit(3).Find(&users).Rows()

	var results []map[string]interface{}
	cols, err := rows.Columns()

	for rows.Next() {
		var row = make([]interface{}, len(cols))
		rows.Scan(row...)

		rowMap := make(map[string]interface{})
		for i, col := range cols {
			fmt.Println(row[i])
			fmt.Println("\n")
			rowMap[col] = row[i]
		}

		results = append(results, rowMap)
	}
	defer rows.Close()

	// send, _ := json.Marshal(results)
	return c.JSON(http.StatusCreated, results)
	// return c.(http.StatusOK, send)
}

func init() {
	cfg, err := ini.Load("conf.ini")
	if err != nil {
		panic(err)
	}
	db.USERNAME = cfg.Section("USERDATABASE").Key("USERNAME").String()
	db.PASSWORD = cfg.Section("USERDATABASE").Key("PASSWORD").String()
	db.HOSTNAME = cfg.Section("USERDATABASE").Key("HOSTNAME").String()
	db.DATABASE = cfg.Section("USERDATABASE").Key("DATABASE").String()
	db.PORT = cfg.Section("USERDATABASE").Key("PORT").String()
}
