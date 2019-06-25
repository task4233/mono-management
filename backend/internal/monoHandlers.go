package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
GetMonos は, itemsテーブルからMonoの情報を集めて返却するメソッドです.
途中でエラーが発生した場合の挙動はwikiを参照してください.


Endpoints


	[GET]	/api/v1/mono
*/
func GetMonos(c *gin.Context) {

	// [TODO]
	// itemsテーブルのidごとに, itemdatasテーブルで該当するデータをくっつける
	// エラー処理はあとで実装
	/*
		{
			SendDefaultHeader(c, "GET")
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  false,
				"message": "JSONの作成に失敗しました.",
			})
			panic(err)
		}
	*/

	SendDefaultHeader(c, "GET")
	resItems := []Item{}
	GetDB().Find(&resItems)
	ms := ResItems{true, resItems}
	c.JSON(http.StatusOK, ms)
}

/*
CreateMono は, リクエストに含まれる情報をitemsテーブルに追加するメソッドです.
途中でエラーが発生した場合の挙動はwikiを参照してください.


Endpoints


	[POST]   /api/v1/mono/new
*/
func CreateMono(c *gin.Context) {
	SendDefaultHeader(c, "POST")

	var reqItem ReqItem
	c.BindJSON(&reqItem)

	CreateDatasByRequest(c, reqItem)

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "作成完了しました",
	})
}

/*
UpdateMonos は, リクエストに該当するitemsテーブルの情報をまとめて更新するメソッドです.
途中でエラーが発生した場合の挙動はwikiを参照してください.


Endpoints


	[PUT]    /api/v1/mono/:monoId
*/
func UpdateMonos(c *gin.Context) {
	// パスの{:monoId}を用いてデータをitemsテーブルから持ってくる
	SendDefaultHeader(c, "PUT")

	var reqItem ReqItem
	c.BindJSON(&reqItem)
	itemID := c.Param("monoId")

	UpdateDatasByRequestAndStrID(c, reqItem, itemID)

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "作成完了しました",
	})
}

/*
DeleteMono は, リクエストに含まれるuserIdを持つデータをitemsテーブルから削除するメソッドです.
途中でエラーが発生した場合の挙動はwikiを参照してください.


Endpoints


	[DELETE] /api/v1/mono/:monoId
*/
func DeleteMono(c *gin.Context) {
	// パスの{:monoId}を用いてデータをitemsテーブルから削除する
	// 関連するデータ群はまとめて削除される
	SendDefaultHeader(c, "DELETE")
	itemID := c.Param("monoId")

	if err := DeleteDatasByStrID(c, itemID); err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "削除完了しました",
	})
}
