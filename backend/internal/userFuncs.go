package internal

import (
	"net/http"

	"crypto/rand"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

/*
GetInfo は, ログイン時に保存されるCookieのtoken情報を元にusersテーブルから該当の情報を返却するメソッドです.
途中でエラーが発生した場合の挙動はwikiを参照してください.


Endpoints


	[GET]	/api/v1/user/info
*/
func GetInfo(c *gin.Context) {

	// Cookieのtoken情報を取得
	// token情報を元にtokensテーブルからuserIdを取得
	// 該当するuserIdを持つデータをusersテーブルから返却
	c.JSON(http.StatusOK, "get info")
}

/*
Login は, userIdとパスワードの正当性が確かめられた時にCookieにtokenをセットして返すメソッドです.
途中でエラーが発生した場合の挙動はwikiを参照してください.


Endpoints


	POST   /api/v1/user/login
*/
func Login(c *gin.Context, name string, password string) {
	// userIdとrawPasswordを取得
	// rawPasswordをhash関数にかけて, usersテーブルからuserIdを元にhashedPassと照合

	// tokenを生成してtokensテーブルにINSERT
	// Cookieにtokenをセットして返す
	db := GetDB()
	user := User{}

	SendDefaultHeader(c, "POST")

	if err := db.Where("name = ?", name).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  false,
			"message": "入力したユーザ名が存在しません",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.HashedPass), []byte(password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  false,
			"message": "パスワードが違います",
		})
		return
	}

	token := Token{}

	randomToken := make([]byte, 64)
	rand.Read(randomToken)

	token.UserId = user.Id
	token.Token = randomToken
	db.Create(&token)

	SetCookie(c, "token", stinrg(randomToken))

	c.JSON(http.StatusOK, gin.H{
		"status": true,
	})
}

/*
Logout はユーザが保持しているCookieを削除するメソッドです.
途中でエラーが発生した場合の挙動はwikiを参照してください.


Endpoints


	[DELETE]   /api/v1/user/logout
*/
func Logout(c *gin.Context) {
	// tokensテーブルからtokenを削除
	// 削除できたらおk
	c.JSON(http.StatusOK, "logout")
}

/*
CreateAccount はリクエストのデータでユーザを新たに作成するメソッドです.
途中でエラーが発生した場合の挙動はwikiを参照してください.


Endpoints


	[POST]   /api/v1/user/new
*/
func CreateAccount(c *gin.Context) {
	// userIdとrawPasswordを取得

	// rawPasswordをハッシュ関数にかけてhashedPassを取得
	// tokensテーブルにhashedPass, userIdを格納
	c.JSON(http.StatusOK, "create account")
}

/*
UpdateAccount はリクエストのデータでユーザ情報を更新するメソッドです.
途中でエラーが発生した場合の挙動はwikiを参照してください.


Endpoints


	[PUT]   /api/v1/user/
*/
func UpdateAccount(c *gin.Context) {
	// usersテーブルからuserIdを元に該当データをupdate

	c.JSON(http.StatusOK, "update account")
}

/*
DeleteAccount は該当アカウントを削除するメソッドです.
途中でエラーが発生した場合の挙動はwikiを参照してください.


Endpoints


	[DELETE]   /api/v1/user/
*/
func DeleteAccount(c *gin.Context) {
	// usersテーブルからuserIdを元に該当データをdelete
	c.JSON(http.StatusOK, "delete account")
}
