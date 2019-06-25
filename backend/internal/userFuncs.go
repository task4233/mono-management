package internal

import (
	"errors"
	"net/http"

	"crypto/rand"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type account struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

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
	SendDefaultHeader(c, "GET")
	db := GetDB()

	user, err := CheckLogin(c)
	if err != nil {
		return
	}

	if err := db.Where("userId = ?", user.Id).First(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "ユーザが見つかりません",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"user":   user,
	})
}

/*
Login は, userIdとパスワードの正当性が確かめられた時にCookieにtokenをセットして返すメソッドです.
途中でエラーが発生した場合の挙動はwikiを参照してください.


Endpoints


	POST   /api/v1/user/login
*/
func Login(c *gin.Context) {
	// userIdとrawPasswordを取得
	// rawPasswordをhash関数にかけて, usersテーブルからuserIdを元にhashedPassと照合

	// tokenを生成してtokensテーブルにINSERT
	// Cookieにtokenをセットして返す
	SendDefaultHeader(c, "POST")
	var json account
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	db := GetDB()
	user := User{}

	if err := db.Where("name = ?", json.Name).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  false,
			"message": "入力したユーザ名が存在しません",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.HashedPass), []byte(json.Password))
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
	token.Token = string(randomToken)
	db.Create(&token)

	SetCookie(c, "token", token.Token)

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
	SendDefaultHeader(c, "DELETE")
	db := GetDB()
	user, err := CheckLogin(c)
	if err != nil {
		return
	}

	token := Token{}
	if err := db.Where("userId = ?", user.Id).First(&token).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "サーバエラー",
		})
		return
	}
	db.Delete(&token)

	c.JSON(http.StatusOK, gin.H{
		"status": true,
	})
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
	SendDefaultHeader(c, "POST")
	var json account
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(), // 仮
		})
	}
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(json.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "サーバエラー",
		})
		return
	}

	user := User{}
	user.Name = json.Name
	user.HashedPass = string(hashedPass)

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

/*
GetUserFromToken はトークンからユーザ情報を取得するメソッドです．
途中でエラーが発生した場合の挙動はwikiを参照してください．

*/
func GetUserFromToken(token string) (User, error) {
	if len(token) == 0 {
		return User{Id: -1}, errors.New("empty token")
	}
	var userToken Token
	if err := GetDB().Where(&Token{Token: token, UserId: 0}).First(&userToken).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return User{Id: -1}, errors.New("login required")
		}
		return User{Id: -1}, errors.New("something went wrong")
	}
	var user User
	if err := GetDB().Where(&User{Id: userToken.UserId, Name: "", HashedPass: ""}).First(&user).Error; err != nil {
		return User{Id: -1}, errors.New("something went wrong")
	}
	return user, nil
}

/*
GetUserFromCookie はCookieからユーザ情報を取得するメソッドです．
途中でエラーが発生した場合の挙動はwikiを参照してください．
*/
func GetUserFromCookie(c *gin.Context) (User, error) {
	token, err := GetCookie(c, "token")
	if err != nil {
		return GetUserFromToken(token)
	} else {
		return User{Id: -1}, errors.New("Login Required")
	}
}

/*
CheckLogin はログイン状態の確認をするメソッドです．
途中でエラーが発生した場合の挙動はwikiを参照してください．
*/
func CheckLogin(c *gin.Context) (User, error) {
	user, err := GetUserFromCookie(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  false,
			"message": err.Error(),
		})
	}
	return user, err
}
