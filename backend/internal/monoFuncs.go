package internal

import (
	"encoding/json"
	"fmt"
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
	// Monoデータをitemsテーブルから取得する
	resItems := []Item{}

	// itemsテーブルの全てのデータを持ってくる
	db := GetDB()
	db.Find(&resItems)

	fmt.Println(resItems)

	// これ無し
	// itemsテーブルのidごとに, itemdatasテーブルで該当するデータをくっつける

	// JSONをまとめてreturnする
	resItemsJSON, err := json.Marshal(resItems)
	if err != nil {
		SendDefaultHeader(c, "GET")
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "JSONの作成に失敗しました.",
		})
		panic(err)
	}

	SendDefaultHeader(c, "GET")
	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   resItemsJSON,
	})

}

/*
CreateMono は, リクエストに含まれる情報をitemsテーブルに追加するメソッドです.
途中でエラーが発生した場合の挙動はwikiを参照してください.


Endpoints


	[POST]   /api/v1/mono/new
*/
func CreateMono(c *gin.Context) {
	// リクエストに含まれる情報のうち, name, userId, tagIdをitemsテーブルに追加

	// リクエストに含まれる情報のうち, datasに含まれるデータ群を1つずつに対して以下の処理を行う
	// + type(primary data), name(データ名)をdatasテーブルに追加する
	// + そのdataId, itemId, 型に該当する値をitemdatasテーブルに追加する
	c.JSON(http.StatusOK, "create mono")
}

/*
UpdateMonos は, リクエストに該当するitemsテーブルの情報をまとめて更新するメソッドです.
途中でエラーが発生した場合の挙動はwikiを参照してください.


Endpoints


	[PUT]    /api/v1/mono/:monoId
*/
func UpdateMonos(c *gin.Context) {
	// パスの{:monoId}を用いてデータをitemsテーブルから持ってくる

	// itemdatasテーブルから{:monoId}を含むデータを全て引き出す
	// リクエストに含まれる情報のうち,
	// keyが, dataIdから取得したdatasテーブルのnameと一致するように
	// itemdatasテーブルのデータを更新する
	c.JSON(http.StatusOK, "update mono")
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

	c.JSON(http.StatusOK, "delete mono")
}
