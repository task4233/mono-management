package internal

import (
	"encoding/base64"
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

	user, err := GetUserFromCookie(c)
	if err != nil {
		return
	}

	if err := db.Where("id = ?", user.ID).First(&user).Error; err != nil {
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
	db := GetDB()

	var json account
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

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

	token := generateToken(user.ID)
	if err := db.Create(&token).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": err.Error(),
		})
	}
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
	user, err := GetUserFromCookie(c)
	if err != nil {
		return
	}

	token := Token{}
	if err := db.Where("userId = ?", user.ID).First(&token).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	if err := db.Delete(&token).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

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
	db := GetDB()

	var json account
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	hashedPass, err := generateHashedPass(json.Password)
	if err != nil {
		createServerErrorMessage(c, "サーバエラー")
		return
	}

	user := User{Name: json.Name, HashedPass: string(hashedPass)}
	if err := db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "既に登録されているユーザ名です",
		})
		return
	}

	token := generateToken(user.ID)
	if err := db.Create(&token).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	SetCookie(c, "token", token.Token)

	c.JSON(http.StatusOK, gin.H{
		"status": true,
	})
}

/*
UpdateAccount はリクエストのデータでユーザ情報を更新するメソッドです.
途中でエラーが発生した場合の挙動はwikiを参照してください.


Endpoints


	[PUT]   /api/v1/user/
*/
func UpdateAccount(c *gin.Context) {
	// usersテーブルからuserIdを元に該当データをupdate
	SendDefaultHeader(c, "PUT")
	db := GetDB()

	var json account
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	user, err := GetUserFromCookie(c)
	if err != nil {
		return
	}

	if len(json.Name) != 0 {
		if err := db.Model(&user).Update("name", json.Name).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if len(json.Password) != 0 {
		hashedPass, err := generateHashedPass(json.Password)
		if err != nil {
			createServerErrorMessage(c, "サーバエラー")
			return
		}
		if err := db.Model(&user).Update("hashedPass", hashedPass).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
	})
}

/*
DeleteAccount は該当アカウントを削除するメソッドです.
途中でエラーが発生した場合の挙動はwikiを参照してください.


Endpoints


	[DELETE]   /api/v1/user/
*/
func DeleteAccount(c *gin.Context) {
	// usersテーブルからuserIdを元に該当データをdelete
	SendDefaultHeader(c, "DELETE")
	db := GetDB()

	user, err := GetUserFromCookie(c)
	if err != nil {
		return
	}

	if err := db.First(&user).Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
	})
}

/*
GetUserFromToken はトークンからユーザ情報を取得するメソッドです．
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
GetUserFromCookie はCookieからユーザ情報を取得するメソッドです．
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

func generateToken(userID int) Token {
	randomToken := make([]byte, 48) // 64文字
	rand.Read(randomToken)
	return Token{Token: base64.URLEncoding.EncodeToString(randomToken), UserID: userID}
}

func createServerErrorMessage(c *gin.Context, message string) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"status":  false,
		"message": message,
	})
}

func generateHashedPass(password string) (string, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPass), err
}
