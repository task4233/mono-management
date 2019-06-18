package internal

import (
	"os"
	"net/http"
	"net/url"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
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
		Name: name,
		Value: url.QueryEscape(val),
		MaxAge: 1892160000,
		Path: "/",
		Secure: false,
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
	db *gorm.DB
	err error
)
func InitDB() {
	DBMS := "mysql"
	USER := os.Getenv("MYSQL_USER")
	PASS := os.Getenv("MYSQL_PASSWORD")
	PROTOCOL := "tcp(db:3306)"
	DBNAME := os.Getenv("MYSQL_DATABASE")
	CONNECT := USER+":"+PASS+"@"+PROTOCOL+"/"+DBNAME+"?charset=utf8mb4&parseTime=True&loc=Local"

	db, err = gorm.Open(DBMS, CONNECT)
	if err != nil {
		panic(err.Error())
	}
}
func GetDB() *gorm.DB {
	return db
}
