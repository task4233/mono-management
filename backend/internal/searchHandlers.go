package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
SearchMonos は, メソッドです.
途中でエラーが発生した場合の挙動はwikiを参照してください.


Endpoints


	[POST]   /api/v1/search/
*/
func SearchMonos(c *gin.Context) {
	var reqSearch ReqSearch
	if err := c.BindJSON(&reqSearch); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "不正なリクエストボディです",
		})
		return
	}

	resItems := []Item{}
	db := GetDB().Begin()
	defer func() {
		if r := recover(); r != nil {
			db.Rollback()
		}
	}()

	if err := db.Where().Find(&resItems).Error; err != nil {
		db.Rollback()
		c.JSON(http.StatusNotFound, gin.H{
			"status":  false,
			"message": "データが存在しません",
		})
		return
	}
	ms := ResItems{true, resItems}
	c.JSON(http.StatusOK, ms)
}
