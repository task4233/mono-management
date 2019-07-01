/*
Package internal は, Ginフレームワークのハンドラの設定等をするパッケージです.

docker-composeの環境下では
	import app/internal
でimportできます.
*/
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

// DB funcs
var (
	db  *gorm.DB
	err error
)

type (
	// User structure
	User struct {
		// ID
		ID int `json:"Id" gorm:"primary_key"`
		// name
		Name string `json:"name"`
		// hashedPass
		HashedPass string `json:"hashedPass" gorm:"column:hashedPass"`
	}

	// Item structure
	Item struct {
		// ID
		ID int `json:"Id" gorm:"primary_key"`
		// name
		Name string `json:"name"`
		// userID
		UserID int `json:"userId" gorm:"column:userId"`
		// tagID
		TagID int `json:"tagId" gorm:"column:tagId"`
	}

	// Itemdata structure
	Itemdata struct {
		// dataID
		DataID int `json:"Id" gorm:"column:dataId; primary_key"`
		// itemID
		ItemID int `json:"itemId" gorm:"column:itemId; primary_key"`
		// num
		Num float64 `json:"num"`
		// str
		Str string `json:"str"`
		// timestamp
		Timestamp *time.Time `json:"timestamp"`
	}

	// Tag structure
	Tag struct {
		// ID
		ID int `json:"Id" gorm:"primary_key"`
		// name
		Name string `json:"name"`
		// parentID
		ParentID int `json:"parentId" gorm:"column:parentId"`
		// userID
		UserID int `json:"userId" gorm:"column:userId"`
	}

	// Data structure
	Data struct {
		// ID
		ID int `json:"Id" gorm:"primary_key"`
		// name
		Name string `json:"name"`
		// type
		Type string `json:"type"`
	}

	// Token structure
	Token struct {
		// token
		Token string `json:"token" gorm:"primary_key"`
		// userID
		UserID int `json:"userId" gorm:"column:userId"`
	}
)

/*
InitDB は, 内部でGorm.Open()を行うメソッドです.

このメソッドはmain関数で一度呼び, Gormの機能を使用したい場合はGetDB()でdbを受け取ってから処理することを推奨しています.
*/
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
	db.LogMode(true)
}

// TableName (users)
func (i *User) TableName() string {
	return "users"
}

// TableName (items)
func (i *Item) TableName() string {
	return "items"
}

// TableName (itemdatas)
func (i *Itemdata) TableName() string {
	return "itemdatas"
}

// TableName (datas)
func (i *Data) TableName() string {
	return "datas"
}

// TableName (tags)
func (i *Tag) TableName() string {
	return "tags"
}

// TableName (tokens)
func (i *Token) TableName() string {
	return "tokens"
}

/*
CloseDB はGORMの接続を閉じる為のメソッドです。

このメソッドはInitDB()と一緒にdeferで呼び出されることが推奨されます。
*/
func CloseDB() {
	GetDB().Close()
}

/*
GetDB は, InitDB()でGorm.Open()したdbのデータのみを受け取るメソッドです.

このメソッドを使用する際はInitDB()を必ず呼び出してください.

以下がサンプルです.
	db := GetDB()

	var users []User
	if err := db.Find(&users).Error; err != nil {
		log.Fatal(err.Error)
	}
*/
func GetDB() *gorm.DB {
	return db
}

/*
SendDefaultHeader は, JSON Headerに情報を付加するメソッドです.

第2引数に許可するHTTPメソッドを渡します.

以下がサンプルです.
	// GETメソッドのみを許可
	SendDefaultHeader(c, "GET")
	// GETおよびPOSTメソッドを許可
	SendDefaultHeader(c, "GET/POST")
*/
func SendDefaultHeader(c *gin.Context, methods string) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", methods)
	c.Header("Access-Control-Max-Age", "86400")
	c.Header("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-HogeApp-Hogeid")
}

/*
SetCookie は, Cookie情報を付加するメソッドです.

第2引数にCookie名を, 第三引数にCookieに付加する値を渡します.

以下がサンプルです.
	// HogeというCookie名で, Fugaという値を保持
	SetCookie(c, "Hoge", "Fuga")

*/
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

/*
GetCookie は, Cookie情報を読み取るメソッドです.

第2引数にCookie名を渡します.

第1返り値に読み取ったCookie情報を, 第2返り値にエラーを返します.

以下がサンプルです.
	// HogeというCookie名の情報を読み取る.
	// 返り値としてcookie_valueというCookie情報と, errというエラーを受け取る.
	cookie_value, err := GetCookie(c, "Hoge")

	// エラーハンドリング
	if err != nil {
		log.Fatal(err.Error)
	}

*/
func GetCookie(c *gin.Context, name string) (string, error) {
	cookie, err := c.Request.Cookie(name)
	if err != nil {
		return "", err
	}
	val, _ := url.QueryUnescape(cookie.Value)
	return val, nil
}
