package internal

import (
	"fmt"
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
	SendDefaultHeader(c, "POST")
	user, err := CheckLogin(c)
	if err != nil {
		return
	}

	var reqSearch ReqSearch
	if err := c.BindJSON(&reqSearch); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error,
		})
		return
	}
	fmt.Println(reqSearch)

	likeStr := "%" + reqSearch.Name + "%"

	resItems := []Item{}

	// ここでLIKE検索をしたい
	// Itemテーブルに対してLIKE検索
	if err := GetDB().Where(&Item{ UserID: user.ID }).Where("name LIKE ?", likeStr).Find(&resItems).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  false,
			"message": "データが存在しません",
		})
		return
	}

	// tag search (include tag children)
	if 0 < reqSearch.TagID {
		var tagItems []Item
		tagChildren, err := GetTagChildren(Tag{ ID: reqSearch.TagID, UserID: user.ID })
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		for _, v := range tagChildren {
			for _, w := range resItems {
				if (v.ID == w.TagID) {
					tagItems = append(tagItems, w)
				}
			}
		}
		resItems = tagItems
	}

	ms := ResItems{true, resItems}
	c.JSON(http.StatusOK, ms)
}
