package internal

import (
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type (
	// users
	User struct {
		// id
		Id int `json:id`
		// name
		Name string `json:name`
		// hashed_pass
		HashedPass string `json:hashedPass`
	}

	// items
	Item struct {
		// id
		Id int `json:id`
		// name
		Name string `json:name`
		// userId
		UserId int `json:userId`
		// tagId
		TagId int `json:tagId`
	}

	// itemdatas
	ItemData struct {
		// dataId
		DataId int `json:dataId`
		// itemId
		ItemId int `json:itemId`
		// num
		Num float64 `json:num`
		// str
		Str string `json:str`
		// timestamp
		Timestamp time.Time `json:timestamp`
	}

	// tags
	Tag struct {
		// id
		Id int `json:id`
		// name
		Name string `json:name`
		// parentId
		ParentId int `json:parentId`
		// userId
		UserId int `json:userId`
	}

	// datas
	Data struct {
		// id
		Id int `json:id`
		// name
		Name string `json:name`
		// type
		Type string `json:type`
	}

	// tokens
	Token struct {
		// token
		Token string `json:token`
		// userId
		UserId int `json:userId`
	}
)

// HTTP funcs
func SendDefaultHeader(c *gin.Context, methods string) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", methods)
	c.Header("Access-Control-Max-Age", "86400")
	c.Header("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

// Cookie funcs
func SetCookie(c *gin.Context, name string, val string) {
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     name,
		Value:    url.QueryEscape(val),
		MaxAge:   1892160000,
		Path:     "/",
		Secure:   false,
		HttpOnly: true,
	})
}

func GetCookie(c *gin.Context, name string) (string, error) {
	cookie, err := c.Request.Cookie(name)
	if err != nil {
		return "", err
	}
	val, _ := url.QueryUnescape(cookie.Value)
	return val, nil
}

// DB funcs
var (
	db  *gorm.DB
	err error
)

func InitDB() {
	DBMS := "mysql"
	USER := os.Getenv("MYSQL_USER")
	PASS := os.Getenv("MYSQL_PASSWORD")
	PROTOCOL := "tcp(db:3306)"
	DBNAME := os.Getenv("MYSQL_DATABASE")
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err = gorm.Open(DBMS, CONNECT)
	if err != nil {
		panic(err.Error())
	}
}
func GetDB() *gorm.DB {
	return db
}
