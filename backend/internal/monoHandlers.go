package internal

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
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
	user, err := CheckLogin(c)
	if err != nil {
		return
	}
	resItems := []Item{}

	if err := GetDB().Where("userId = ?", user.ID).Find(&resItems).Error; err != nil && !gorm.IsRecordNotFoundError(err) {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  false,
			"message": "何らかのエラーが発生しました",
		})
		return
	}
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
	if err := c.BindJSON(&reqItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "不正なリクエストボディです",
		})
		return
	}

	if err := CreateDatasByRequest(c, reqItem); err != nil {
		return
	}

	ReturnStatusOKWithStrMessage(c, "作成完了しました")

}

/*
GETMono は, mono単体の全情報を取得する関数です.
途中でエラーが発生した場合の挙動はwikiを参照してください.


Endpoints


	[GET]    /api/v1/mono/:monoId
*/
func GetMonoData(c *gin.Context) {
	SendDefaultHeader(c, "GET")
	user, err := CheckLogin(c)
	if err != nil {
		return
	}
	monoId, err := strconv.Atoi(c.Param("monoId"))
	if err != nil || monoId < 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "リクエストエラー",
		})
		return
	}
	db := GetDB()
	mono := Item{ID: monoId, UserID: user.ID}
	if err := db.Where(&mono).First(&mono).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  false,
				"message": "指定されたmonoが見つかりませんでした",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "不明なエラー",
		})
		return
	}
	/*var monodata []Itemdata
	if err := db.Where(&Itemdata{ ItemID: mono.ID }).Find(&monodata).Error; err != nil && !gorm.IsRecordNotFoundError(err) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"message": "不明なエラー",
		})
	}*/
	var jid []JoinedItemData
	if err := GetJoinedItemData().Where("itemId = ?", mono.ID).Find(&jid).Error; err != nil && !gorm.IsRecordNotFoundError(err) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "不明なエラー",
		})
		return
	}
	retData := gin.H{}
	for _, v := range jid {
		retData[v.Name] = v
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"item":   mono,
		"data":   retData,
	})
	return
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

	if err := UpdateDatasByRequestAndStrID(c, reqItem, itemID); err != nil {
		return
	}
	ReturnStatusOKWithStrMessage(c, "更新完了しました")
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
	ReturnStatusOKWithStrMessage(c, "削除完了しました")
}
