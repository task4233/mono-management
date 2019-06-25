package internal

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// Res is the interface of Item structure
type Res Item
type ResItems struct {
	Status bool
	Item   []Item
}

/*
AddStatusMessageForItems は与えられたjsonデータにステータスを追加して返すメソッドです.
*/
func (i Item) AddStatusMessageForItems() ([]byte, error) {
	return json.Marshal(struct {
		Res
		Status bool
	}{
		Res:    Res(i),
		Status: true,
	})
}

// Res is the interface of Item structure
type ResItem struct {
	Status bool
	Item   Item
}

/*
AddStatusMessageForItem は与えられたjsonデータにステータスを追加して返すメソッドです.
*/
func (i Item) AddStatusMessageForItem() ([]byte, error) {
	return json.Marshal(struct {
		Res
		Status bool
	}{
		Res:    Res(i),
		Status: true,
	})
}

func CreateDatasByRequest(c *gin.Context, reqItem ReqItem) error {
	// createTokenができるまでスタブにしとく
	// 認証
	/*
		reqUser, err := CheckLogin(c)
		if err != nil {
			return
		}
		fmt.Printf("%+v\n", reqUser)
	*/
	// スタブ
	reqUser := User{Id: 1}

	// fmt.Printf("%+v\n", reqItem) // for debug

	newItem := Item{Name: reqItem.Name, Userid: reqUser.Id, Tagid: reqItem.TagId}
	// fmt.Printf("%+v\n", newItem)  // for debug

	db := GetDB()
	// トランザクション開始
	tx := db.Begin()
	db.Create(&newItem)
	if tx.Error != nil {
		return tx.Error
	}

	for _, data := range reqItem.Datas {
		newData := Data{Name: data.Name, Type: data.Type}
		db.Create(&newData)
		// fmt.Printf("%+v\n", newData) // for debug

		switch data.Type {
		case "num":
			numVal, err := strconv.ParseFloat(data.Value, 64)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"status":  false,
					"message": "型が数字ではありません",
				})
				return err
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
				return err
			}
			db.Create(&Itemdata{DataId: newItem.Id, ItemId: newItem.Id, Timestamp: &timestampValue})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  false,
				"message": "無効な型です",
			})
			return err
		}
	}
	return nil
}

func UpdateDatasByRequestAndStrID(c *gin.Context, reqItem ReqItem, itemID string) error {
	// createTokenができるまでスタブにしとく
	// 認証
	/*
		reqUser, err := CheckLogin(c)
		if err != nil {
			return
		}
		fmt.Printf("%+v\n", reqUser)
	*/
	// スタブ
	reqUser := User{Id: 1}

	// fmt.Printf("%+v\n", reqItem) // for debug
	// fmt.Printf("%+v\n", newItem)  // for debug

	editItem := Item{}
	if editItem.Id, err = strconv.Atoi(itemID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  false,
			"message": "不正なidです",
		})
		return err
	}

	db := GetDB()
	db.Model(&editItem).Updates(Item{Name: reqItem.Name, Userid: reqUser.Id, Tagid: reqItem.TagId})

	editItemDatas := []Itemdata{}
	db.Where(&Itemdata{ItemId: editItem.Id}).Find(&editItemDatas)

	for _, editItemData := range editItemDatas {
		// fmt.Printf("[%d]%+v\n", index, delItemdata) // for debug

		// datasテーブルのそのレコードをupdate
		for _, data := range reqItem.Datas {
			editData := Data{Id: editItemData.DataId}
			db.First(&editData)
			if editData.Name != data.Name {
				continue
			}

			switch data.Type {
			case "num":
				numVal, err := strconv.ParseFloat(data.Value, 64)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{
						"status":  false,
						"message": "型が数字ではありません",
					})
					return err
				}
				db.Model(&editItemData).Updates(&Itemdata{DataId: editData.Id, ItemId: editItem.Id, Num: numVal, Timestamp: nil})
			case "str":
				db.Model(&editItemData).Updates(&Itemdata{DataId: editData.Id, ItemId: editItem.Id, Str: data.Value, Timestamp: nil})
			case "timestamp":
				// time.Parse("layout", "value")
				timestampValue, err := time.Parse(time.RFC1123, data.Value)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{
						"status":  false,
						"message": "型が日付ではありません",
					})
					return err
				}
				db.Model(&editItemData).Updates(&Itemdata{DataId: editItem.Id, ItemId: editItem.Id, Timestamp: &timestampValue})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{
					"status":  false,
					"message": "無効な型です",
				})
				return err
			}
		}
	}

	return nil
}

func DeleteDatasByStrID(c *gin.Context, itemId string) error {
	delItem := Item{}
	if delItem.Id, err = strconv.Atoi(itemId); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  false,
			"message": "不正なidです",
		})
		return err
	}
	// fmt.Printf("%+v\n", delItem) // for debug

	db := GetDB()
	db.First(&delItem)

	delItemdatas := []Itemdata{}
	db.Where(&Itemdata{ItemId: delItem.Id}).Find(&delItemdatas)

	for _, delItemdata := range delItemdatas {
		// fmt.Printf("[%d]%+v\n", index, delItemdata) // for debug

		// datasテーブルのそのレコードをdelete
		delData := Data{Id: delItemdata.DataId}
		db.First(&delData).Delete(&delData)

		// 保持しておいたdataidとitemidを持つレコードをitemdatasからdelete
		db.Delete(&delItemdata)
	}

	// 最後にitemを削除
	db.Delete(&delItem)
	return nil
}

func ReturnStatusOKWithStrMessage(c *gin.Context, message string) {
	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": message,
	})
}

//
// Table Setting
//

//
// received request form
//
type ReqItemData struct {
	Name  string
	Value string
	Type  string
}

type ReqItem struct {
	Name  string
	TagId int `json:"tagId"`
	Datas []ReqItemData
}
