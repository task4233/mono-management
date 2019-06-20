package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
GetTags はuserIdが一致するTagを全て取得するメソッドです.
途中でエラーが発生した場合の挙動はwikiを参照してください.


Endpoints


	[GET]    /api/v1/tag/
*/
func GetTags(c *gin.Context) {
	//
	c.JSON(http.StatusOK, "get tags")
}

/*
CreateTag はuserIdを付加したTagを全て取得するメソッドです.
途中でエラーが発生した場合の挙動はwikiを参照してください.


Endpoints


	[POST]   /api/v1/tag/new
*/
func CreateTag(c *gin.Context) {
	c.JSON(http.StatusOK, "crate tag")
}

/*
UpdateTag はuserIdを付加したTagを更新するメソッドです.
途中でエラーが発生した場合の挙動はwikiを参照してください.


Endpoints


	[PUT]    /api/v1/tag/:tagId
*/
func UpdateTag(c *gin.Context) {
	// samp
	c.JSON(http.StatusOK, "update tag")
}

/*
DeleteTag tagIdを元にデータを削除するメソッドです.
途中でエラーが発生した場合の挙動はwikiを参照してください.


Endpoints


	[DELETE] /api/v1/tag/:tagId
*/
func DeleteTag(c *gin.Context) {
	// samp
	c.JSON(http.StatusOK, "delete tag")
}
