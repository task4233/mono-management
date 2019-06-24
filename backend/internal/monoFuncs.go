package internal

import (
	"fmt"
	"net/http"
	"strconv"

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
	// [TODO]
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
	ms := ResItems{true, resItems}

	SendDefaultHeader(c, "GET/POST")

	c.JSON(http.StatusOK, ms)

}

/*
CreateMono は, リクエストに含まれる情報をitemsテーブルに追加するメソッドです.
途中でエラーが発生した場合の挙動はwikiを参照してください.


Endpoints


	[POST]   /api/v1/mono/new
*/
func CreateMono(c *gin.Context) {
	// リクエストに含まれる情報のうち, name, userId, tagIdをitemsテーブルに追加
	// 受け取るjsonの形式は以下の通り
	//```json
	// +name
	// +userId
	// +tagId
	//```

	var reqItem Item
	c.BindJSON(&reqItem)
	db := GetDB()
	fmt.Println(reqItem)
	db.Create(&reqItem)
	if db.NewRecord(reqItem) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "データを作成できませんでした",
		})
	} else {
		fmt.Println(reqItem)

		// リクエストに含まれる情報のうち, datasに含まれるデータ群を1つずつに対して以下の処理を行う
		// + type(primary data), name(データ名)をdatasテーブルに追加する
		// + そのdataId, itemId, 型に該当する値をitemdatasテーブルに追加する
		ms := ResItem{true, reqItem}
		c.JSON(http.StatusOK, ms)
	}
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
	SendDefaultHeader(c, "DELETE")
	delItem := Item{}
	itemID := c.Param("monoId")
	delItem.Id, err = strconv.Atoi(itemID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "削除対象のmonoが見つかりませんでした",
		})
		panic(err)
	}

	db := GetDB()
	Itemdatas := []Itemdata{}
	db.Where("Itemid=?", itemID).Find(&Itemdatas)
	fmt.Printf("%+v\n", Itemdatas)

	db.First(&delItem)

	deldata := Data{}
	deldata.Id = delItem.Id

	db.Delete(&delItem)

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "削除完了しました",
	})
}
