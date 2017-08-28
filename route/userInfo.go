package route

import (
	"bytes"
	"net/http"

	"../model"
	"github.com/go-ini/ini"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //mysql driver
	"github.com/labstack/echo"
)

var (
	dbLink string
)

// GetUserInfo is router for GetUserInfo
func GetUserInfo(c echo.Context) error {
	var user model.User
	id := c.Param("id")

	db, err := gorm.Open("mysql", dbLink)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.LogMode(true)
	query := db.Table("user").First(&user, "id = ?", id).Model(&user).Error
	if query != nil {
		panic(query)
	}
	return c.JSON(http.StatusCreated, &user)
}

// PutUserInfo is put userinfo
func PutUserInfo(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusCreated, id)
}

func init() {
	cfg, err := ini.Load("conf.ini")
	if err != nil {
		panic(err)
	}
	USERNAME := cfg.Section("USERDATABASE").Key("USERNAME").String()
	PASSWORD := cfg.Section("USERDATABASE").Key("PASSWORD").String()
	HOSTNAME := cfg.Section("USERDATABASE").Key("HOSTNAME").String()
	DATABASE := cfg.Section("USERDATABASE").Key("DATABASE").String()
	PORT := cfg.Section("USERDATABASE").Key("PORT").String()

	b := bytes.Buffer{}
	b.WriteString(USERNAME)
	b.WriteString(":")
	b.WriteString(PASSWORD)
	b.WriteString("@tcp(")
	b.WriteString(HOSTNAME)
	b.WriteString(":")
	b.WriteString(PORT)
	b.WriteString(")/")
	b.WriteString(DATABASE)
	b.WriteString("?charset=utf8&parseTime=True")
	dbLink = b.String()
}
