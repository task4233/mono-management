package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Res is the interface of Item structure
type Res Item
type ResItems struct {
	Status bool
	Data   []Item
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
	Data   Item
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
	reqUser, err := CheckLogin(c)
	if err != nil {
		return errors.New("Login error")
	}
	// fmt.Printf("%+v\n", reqUser) // for debug

	// fmt.Printf("%+v\n", reqItem) // for debug

	newItem := Item{Name: reqItem.Name, UserID: reqUser.ID, TagID: reqItem.TagID}
	fmt.Printf("%+v\n", newItem) // for debug

	// トランザクション開始
	db := GetDB().Begin()
	defer func() {
		if r := recover(); r != nil {
			db.Rollback()
		}
	}()

	currentTag := Tag{ID: newItem.TagID}
	if err := db.First(&currentTag).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "タグが存在していません",
		})
		db.Rollback()
		return err
	}

	// タグを保有していないユーザ
	if currentTag.UserID != newItem.UserID {
		CreateServerErrorMessage(c, "タグを保有していないユーザです")
		db.Rollback()
		return errors.New("タグを保有していないユーザです")
	}

	if err := db.Create(&newItem).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "tagIdが不正です",
		})
		db.Rollback()
		return err
	}

	for _, data := range reqItem.Data {
		newData := Data{Name: data.Name, Type: data.Type}
		if err := db.Create(&newData).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  false,
				"message": "データ名が不正です",
			})
			db.Rollback()
			return err
		}
		// fmt.Printf("%+v\n", newData) // for debug

		switch data.Type {
		case "num":
			numVal, err := strconv.ParseFloat(data.Value, 64)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": "型が数字ではありません",
				})
				return err
			}
			if err := db.Create(&Itemdata{DataID: newData.ID, ItemID: newItem.ID, Num: numVal, Timestamp: nil}).Error; err != nil {
				db.Rollback()
				c.JSON(http.StatusInternalServerError, gin.H{
					"status":  false,
					"message": "データ名が不正です",
				})
				return err
			}
		case "str":
			if err := db.Create(&Itemdata{DataID: newData.ID, ItemID: newItem.ID, Str: data.Value, Timestamp: nil}).Error; err != nil {
				db.Rollback()
				c.JSON(http.StatusInternalServerError, gin.H{
					"status":  false,
					"message": "データ名が不正です",
				})
				return err
			}
		case "timestamp":
			// time.Parse("layout", "value")
			timestampValue, err := time.Parse(time.RFC1123, data.Value)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": "型が日付ではありません",
				})
				return err
			}

			if err := db.Create(&Itemdata{DataID: newData.ID, ItemID: newItem.ID, Timestamp: &timestampValue}).Error; err != nil {
				db.Rollback()
				c.JSON(http.StatusInternalServerError, gin.H{
					"status":  false,
					"message": "timestampに関するデータが不正です",
				})
				return err
			}
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  false,
				"message": "無効な型です",
			})
			return errors.New("Invalid Structure")
		}
	}
	if err := db.Commit().Error; err != nil {
		db.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "リクエストされたデータに不備があり, データベースにコミットできません",
		})
		return err
	}
	return nil
}

func UpdateDatasByRequestAndStrID(c *gin.Context, reqItem ReqItem, itemID string) error {
	reqUser, err := CheckLogin(c)
	if err != nil {
		return errors.New("Login error")
	}

	editItem := Item{}
	if editItem.ID, err = strconv.Atoi(itemID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  false,
			"message": "不正なIDです",
		})
		return err
	}

	editItem.UserID = reqUser.ID

	db := GetDB().Begin()
	defer func() {
		if r := recover(); r != nil {
			db.Rollback()
		}
	}()

	fmt.Printf("ok")
	if err := db.Where(&editItem).First(&editItem).Error; err != nil {
		db.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "そのようなアイテムは存在しません",
		})
		return err
	}
	fmt.Printf("ok2")
	if err := db.Model(&editItem).Updates(Item{Name: reqItem.Name, UserID: reqUser.ID, TagID: reqItem.TagID}).Error; err != nil {
		db.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "データ名が不正です",
		})
		return err
	}

	fmt.Printf("ok3")
	editItemDatas := []Itemdata{}
	if err := db.Where(&Itemdata{ItemID: editItem.ID}).Find(&editItemDatas).Error; err != nil {
		db.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "データ名が不正です",
		})
		return err
	}

	fmt.Printf("ok4")
	for _, editItemData := range editItemDatas {
		// fmt.Printf("[%d]%+v\n", index, delItemdata) // for debug

		// datasテーブルのそのレコードをupdate
		for _, data := range reqItem.Data {
			editData := Data{ID: editItemData.DataID}
			if err := db.First(&editData).Error; err != nil {
				db.Rollback()
				c.JSON(http.StatusInternalServerError, gin.H{
					"status":  false,
					"message": "データ名が不正です",
				})
				return err
			}
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
				if err := db.Model(&editItemData).Updates(&Itemdata{DataID: editData.ID, ItemID: editItem.ID, Num: numVal, Timestamp: nil}).Error; err != nil {
					db.Rollback()
					c.JSON(http.StatusInternalServerError, gin.H{
						"status":  false,
						"message": "numのデータが不正です",
					})
					return err
				}
			case "str":
				if err := db.Model(&editItemData).Updates(&Itemdata{DataID: editData.ID, ItemID: editItem.ID, Str: data.Value, Timestamp: nil}).Error; err != nil {
					db.Rollback()
					c.JSON(http.StatusInternalServerError, gin.H{
						"status":  false,
						"message": "strのデータが不正です",
					})
					return err
				}
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
				if err := db.Model(&editItemData).Updates(&Itemdata{DataID: editData.ID, ItemID: editItem.ID, Timestamp: &timestampValue}).Error; err != nil {
					db.Rollback()
					c.JSON(http.StatusInternalServerError, gin.H{
						"status":  false,
						"message": "timestampのデータが不正です",
					})
					return err
				}
			default:
				c.JSON(http.StatusInternalServerError, gin.H{
					"status":  false,
					"message": "無効な型です",
				})
				return err
			}
		}
	}
	fmt.Printf("ok6")
	if err := db.Commit().Error; err != nil {
		db.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "リクエストされたデータに不備があり, データベースにコミットできません",
		})
		return err
	}

	return nil
}

func DeleteDatasByStrID(c *gin.Context, itemID string) error {
	reqUser, err := CheckLogin(c)
	if err != nil {
		return errors.New("Login error")
	}
	delItem := Item{}
	if delItem.ID, err = strconv.Atoi(itemID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  false,
			"message": "不正なIDです",
		})
		return err
	}
	delItem.UserID = reqUser.ID
	// fmt.Printf("%+v\n", delItem) // for debug

	db := GetDB().Begin()
	defer func() {
		if r := recover(); r != nil {
			db.Rollback()
		}
	}()

	if err := db.Where(&delItem).First(&delItem).Error; err != nil {
		db.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "そのようなアイテムは存在しません",
		})
		return err
	}

	delItemdatas := []Itemdata{}
	if err := db.Where(&Itemdata{ItemID: delItem.ID}).Find(&delItemdatas).Error; err != nil {
		db.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "削除しようとしているデータが見つかりません",
		})
		return err
	}

	for _, delItemdata := range delItemdatas {
		// fmt.Printf("[%d]%+v\n", index, delItemdata) // for debug

		// datasテーブルのそのレコードをdelete
		delData := Data{ID: delItemdata.DataID}
		if err := db.First(&delData).Delete(&delData).Error; err != nil {
			db.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  false,
				"message": "データ名が不正です",
			})
			return err
		}

		// 保持しておいたdataIDとitemIDを持つレコードをitemdatasからdelete
		if err := db.Delete(&delItemdata).Error; err != nil {
			db.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  false,
				"message": "データが不正です",
			})
			return err
		}
	}

	// 最後にitemを削除
	if err := db.Delete(&delItem).Error; err != nil {
		db.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "データが削除できません",
		})
		return err
	}
	if err := db.Commit().Error; err != nil {
		db.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "リクエストされたデータに不備があり, データベースにコミットできません",
		})
		return err
	}
	return nil
}

func ReturnStatusOKWithStrMessage(c *gin.Context, message string) {
	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": message,
	})
}

func GetJoinedItemData() *gorm.DB {
	return GetDB().Table("itemdatas").Select("*").Joins("left join datas on itemdatas.dataId = datas.id")
}

//
// Table Setting
//

//
// received request form
//
type ReqItemData struct {
	Name  string `json: "name"`
	Value string `json: "value"`
	Type  string `json: "type"`
}

type ReqItem struct {
	Name  string        `json: "name"`
	TagID int           `json:"tagId" gorm:"tagId"`
	Data  []ReqItemData `json: "data"`
}

type JoinedItemData struct {
	Itemdata
	Data
}
