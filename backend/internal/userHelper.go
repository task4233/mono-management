package internal

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

/*
Account はアカウント作成・更新時にリクエストボディとして送信されるべきユーザ情報です
*/
type Account struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

/*
GetUserFromToken はトークンからユーザ情報を取得する関数です．
途中でエラーが発生した場合の挙動はwikiを参照してください．

*/
func GetUserFromToken(token string) (User, error) {
	if len(token) == 0 {
		return User{ID: -1}, errors.New("ログインしてください")
	}
	var userToken Token
	if err := GetDB().Where(&Token{Token: token, UserID: 0}).First(&userToken).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return User{ID: -1}, errors.New("サーバエラー")
		}
		return User{ID: -1}, errors.New("サーバエラー")
	}
	var user User
	if err := GetDB().Where(&User{ID: userToken.UserID, Name: "", HashedPass: ""}).First(&user).Error; err != nil {
		return User{ID: -1}, errors.New("サーバエラー")
	}
	return user, nil
}

/*
GetUserFromCookie はCookieからユーザ情報を取得する関数です．
途中でエラーが発生した場合の挙動はwikiを参照してください．
*/
func GetUserFromCookie(c *gin.Context) (User, error) {
	token, err := GetCookie(c, "token")
	if err != nil {
		return User{ID: -1}, errors.New("ログインしてください")
	}
	return GetUserFromToken(token)
}

/*
CheckLogin はログイン状態の確認をする関数です．
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

/*
CreateServerErrorMessage はInternalServerErrorをメッセージ付きで作成する関数です．
*/
func CreateServerErrorMessage(c *gin.Context, message string) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"status":  false,
		"message": message,
	})
}
