package internal

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

/*
GetMonos は, itemsテーブルからMonoの情報を集めて返却するメソッドです.
途中でエラーが発生した場合の挙動はwikiを参照してください.


Endpoints


	[GET]	/api/v1/mono
*/
func GetMonos(c *gin.Context) {
	SendDefaultHeader(c, "GET")

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

	c.JSON(http.StatusOK, ms)

}

/*
CreateMono は, リクエストに含まれる情報をitemsテーブルに追加するメソッドです.
途中でエラーが発生した場合の挙動はwikiを参照してください.


Endpoints


	[POST]   /api/v1/mono/new
*/
func CreateMono(c *gin.Context) {
	// リクエストを取得
	// リクエストに含まれる情報のうち, name, userId, tagIdをitemsテーブルに追加
	// 受け取るjsonの形式は以下の通り
	//```json
	// +name
	// +userId
	// +tagId
	//```

	SendDefaultHeader(c, "POST")

	// createTokenができるまでスタブにしとく
	// 認証
	/*
		reqUser, err := CheckLogin(c)
		if err != nil {
			return
		}
		fmt.Printf("%+v\n", reqUser)
	*/

	reqUser := User{}
	reqUser.Id = 1
	// リクエストを取得
	// リクエストの形式は以下のwikiみて
	//
	var reqItem ReqItem
	c.BindJSON(&reqItem)
	fmt.Printf("%+v\n", reqItem)

	// リクエストのuserId, tagId, nameをitemsに追加
	newItem := Item{}
	newItem.Name = reqItem.Name
	newItem.Userid = reqUser.Id
	newItem.Tagid = reqItem.TagId

	fmt.Printf("%+v\n", newItem)

	db := GetDB()
	db.Create(&newItem)

	// リクエストボディをfor文で回す
	for _, data := range reqItem.Datas {
		// datasの"name"にリクエストボディの"name"を格納
		// datasの"vtype"に型を判断してリクエストボディの"type"を格納

		newData := Data{Name: data.Name, Type: data.Type}
		db.Create(&newData)
		fmt.Printf("%+v\n", newData)

		// itemdatasに"value"の値を格納してcreate
		switch data.Type {
		case "num":
			numVal, err := strconv.ParseFloat(data.Value, 64)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"status":  false,
					"message": "型が数字ではありません",
				})
				return
			}
			db.Create(&Itemdata{DataId: newData.Id, ItemId: newItem.Id, Num: numVal, Timestamp: nil})
		case "str":
			db.Create(&Itemdata{DataId: newData.Id, ItemId: newItem.Id, Str: data.Value, Timestamp: nil})
		case "timestamp":
			// time.Parse("layout", "value")
			timestampValue, err := time.Parse(time.RFC1123, data.Value)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"status":  false,
					"message": "型が日付ではありません",
				})
				return
			}
			db.Create(&Itemdata{DataId: newData.Id, ItemId: newItem.Id, Timestamp: &timestampValue})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  false,
				"message": "無効な型です",
			})
			return
		}

	}

	// ここでfor文おわり

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
	updateItemdata := Itemdata{}
	itemID := c.Param("monoId")
	updateItemdata.ItemId, err = strconv.Atoi(itemID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "削除対象のmonoが見つかりませんでした",
		})
		panic(err)
	}

	db := GetDB()
	db.Where("Itemid=?", itemID).Find(&updateItemdata)
	fmt.Printf("%+v\n", updateItemdata)

	db.Find(&updateItemdata)

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "更新完了しました",
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

	// itemidの保持(itemID)
	itemID := c.Param("monoId")

	// idでItemを取得
	delItem := Item{}
	delItem.Id, err = strconv.Atoi(itemID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  false,
			"message": "削除対象のmonoが見つかりませんでした",
		})
		panic(err)
	}
	fmt.Printf("%+v\n", delItem)

	db := GetDB()
	db.First(&delItem)

	// itemdatasを全てもっtくる(where itemid = Id)
	delItemdatas := []Itemdata{}
	db.Where("Itemid=?", delItem.Id).Find(&delItemdatas)

	// ここからforループ
	for index, delItemdata := range delItemdatas {
		fmt.Printf("[%d]%+v\n", index, delItemdata)
		// for文で1つ1つみて, dataidを元に参照したdatasのnameが同じなら

		// datasテーブルのそのレコードをdelete
		delData := Data{}
		delData.Id = delItemdata.DataId
		db.First(&delData)
		db.Delete(&delData)

		// 保持しておいたdataidとitemidを持つレコードをitemdatasからdelete
		db.Delete(&delItemdata)
	}
	// ここでforループ終わり

	// 最後にitemを削除
	db.Delete(&delItem)

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "削除完了しました",
	})
}
