package internal

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

/*
GetTags はuserIdが一致するTagを全て取得するメソッドです.
途中でエラーが発生した場合の挙動はwikiを参照してください.


Endpoints


	[GET]    /api/v1/tag/
*/
func GetTagsHandler(c *gin.Context) {
	SendDefaultHeader(c, "GET")
	user, err := CheckLogin(c)
	if err != nil {
		return
	}
	tags, err := GetTags(0, user.ID)
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"message": "何らかのエラーが発生しました",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"tags": tags,
	})
}

/*
CreateTag はuserIdを付加したTagを全て取得するメソッドです.
途中でエラーが発生した場合の挙動はwikiを参照してください.


Endpoints


	[POST]   /api/v1/tag/new
*/
func CreateTagHandler(c *gin.Context) {
	SendDefaultHeader(c, "POST")
	user, err := CheckLogin(c)
	if err != nil {
		return
	}
	var newTag Tag
	if err := c.ShouldBindJSON(&newTag); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"message": "リクエストエラー",
		})
		return
	}
	if newTag.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"message": "タグ名を空にすることはできません。",
		})
		return
	}
	newTag.UserID = user.ID
	crtedTag, err := CreateTag(newTag)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"message": "何らかのエラーが発生しました",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"tag": crtedTag,
	})
}

/*
UpdateTag はuserIdを付加したTagを更新するメソッドです.
途中でエラーが発生した場合の挙動はwikiを参照してください.


Endpoints


	[PUT]    /api/v1/tag/:tagId
*/
func UpdateTagHandler(c *gin.Context) {
	SendDefaultHeader(c, "PUT")
	user, err := CheckLogin(c)
	if err != nil {
		return
	}
	var updTag Tag
	if err := c.ShouldBindJSON(&updTag); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"message": "リクエストエラー",
		})
		return
	}
	tagid, err := strconv.Atoi(c.Param("tagId"))
	if err != nil || tagid < 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"message": "リクエストエラー",
		})
		return
	}
	updTag.ID = tagid
	updTag.UserID = user.ID
	if _, err := GetTagParents(updTag); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"message": err.Error(),
		})
		return
	}
	updedTag, err := UpdateTag(Tag{ ID: tagid, UserID: user.ID }, updTag)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			c.JSON(http.StatusNotFound, gin.H{
				"status": false,
				"message": "指定されたタグが見つかりませんでした",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": false,
				"message": "何らかのエラーが発生しました",
			})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"tag": updedTag,
	})
}

/*
DeleteTag tagIdを元にデータを削除するメソッドです.
途中でエラーが発生した場合の挙動はwikiを参照してください.


Endpoints


	[DELETE] /api/v1/tag/:tagId
*/
func DeleteTagHandler(c *gin.Context) {
	SendDefaultHeader(c, "DELETE")
	user, err := CheckLogin(c)
	if err != nil {
		return
	}
	var delTag Tag
	tagid, err := strconv.Atoi(c.Param("tagId"))
	if err != nil || tagid < 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"message": "リクエストエラー",
		})
		return
	}
	delTag.ID = tagid
	delTag.UserID = user.ID
	deltedTag, err := DeleteTag(Tag{ ID: tagid, UserID: user.ID }, delTag)

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			c.JSON(http.StatusNotFound, gin.H{
				"status": false,
				"message": "指定されたタグが見つかりませんでした",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": false,
				"message": "何らかのエラーが発生しました。親タグを消すには、先に小タグを消去してください。",
			})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"tag": deltedTag,
	})
}
