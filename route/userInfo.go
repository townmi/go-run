package route

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
	"time"

	"../model"
	"github.com/go-ini/ini"
	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //mysql driver
	"github.com/labstack/echo"
)

var (
	userInfoDbLink  string
	graphqlUserData map[string]model.User
)

// Response is
type Response struct {
	Code       int         `json:"code"`
	Data       *model.User `json:"data"`
	ServerTime int64       `json:"serverTime"`
}

var venuesType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"displayName": &graphql.Field{
				Type: graphql.String,
			},
			"password": &graphql.Field{
				Type: graphql.String,
			},
			"role": &graphql.Field{
				Type: graphql.String,
			},
			"email": &graphql.Field{
				Type: graphql.String,
			},
			"mobile": &graphql.Field{
				Type: graphql.String,
			},
			"locationLat": &graphql.Field{
				Type: graphql.String,
			},
			"locationLon": &graphql.Field{
				Type: graphql.String,
			},
			"subscribe": &graphql.Field{
				Type: graphql.Int,
			},
			"cityId": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"user": &graphql.Field{
				Type: venuesType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					db, err := gorm.Open("mysql", userInfoDbLink)
					if err != nil {
						panic(err)
					}
					defer db.Close()
					db.LogMode(true)
					var user model.User
					// query := db.Table("user").Where(&model.User{ID: p.Source.(model.User).ID}).Model(&user).Error
					query := db.Table("user").First(&user, "id = ?", p.Args["id"]).Model(&user).Error
					if query != nil {
						return nil, nil
					}
					return user, nil
				},
			},
		},
	})

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: queryType,
	},
)

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}

// UsersGraphql is
func UsersGraphql(c echo.Context) error {
	result := executeQuery(c.QueryParam("query"), schema)
	return c.JSONPretty(http.StatusOK, &result, "  ")
}

// GetUserInfo is router for GetUserInfo
func GetUserInfo(c echo.Context) error {
	var (
		user model.User
		res  Response
	)
	t := time.Now()
	res.ServerTime = t.Unix() * 1000
	id := c.Param("id")
	tableType := c.QueryParam("type")
	// ids := c.QueryParam("ids")
	// if tableType {

	// }
	db, err := gorm.Open("mysql", userInfoDbLink)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.LogMode(true)
	query := db.Table(strings.ToUpper(tableType)).First(&user, "id = ?", id).Model(&user).Error
	if query != nil {
		res.Code = 404
		return c.JSONPretty(http.StatusOK, &res, "  ")
	}
	res.Code = 200
	res.Data = &user
	return c.JSONPretty(http.StatusOK, &res, "  ")
}

// PutUserInfo is put userinfo
func PutUserInfo(c echo.Context) error {
	id := c.Param("id")
	return c.JSONPretty(http.StatusOK, id, "	")
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
	userInfoDbLink = b.String()
}
