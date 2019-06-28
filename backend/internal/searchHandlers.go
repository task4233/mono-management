package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
SearchMonos は, POSTされたmonoのキーワードをもとに該当するデータを全て返却するメソッドです.
途中でエラーが発生した場合の挙動はwikiを参照してください.


Endpoints


	[POST]   /api/v1/search/
*/
func SearchMonos(c *gin.Context) {
	var reqSearch ReqSearch
	if err := c.BindJSON(&reqSearch); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error,
		})
		return
	}

    likeStr := "%" + reqSearch.Name + "%"
    
	resItems := []Item{}
	db := GetDB().Begin()
	defer func() {
		if r := recover(); r != nil {
			db.Rollback()
		}
	}()

    // ここでLIKE検索をしたい
    // Itemテーブルに対してLIKE検索
	if err := db.Where(&Item{TagID: reqSearch.TagID}).Where("name LIKE ?", likeStr).Find(&resItems).Error; err != nil {
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
