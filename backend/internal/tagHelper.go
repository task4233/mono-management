package internal

import(
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

/*
GetTags はタグIDとユーザIDを引数にとり、tagsテーブルから該当するタグのスライスを取得する関数です。

第一引数にタグID、第二引数にユーザIDを取り、 ([]Tag, error) を返り値とします。

GORMの仕様に従い、引数に0を指定した場合には、IDを指定せずに検索を行います。例えば、 GetTags(0, 3) とした場合には、ユーザIDが3の全てのタグを取得します。

*/
func GetTags(tagId int, userId int) ([]Tag, error) {
	var tags []Tag
	err := GetDB().Where(&Tag{ ID: tagId, Name: "", ParentID: 0, UserID: userId }).Find(&tags).Error
	return tags, err
}

func CreateTag(tag Tag) (Tag, error) {
	err := GetDB().Create(&tag).Error
	return tag, err
}

func UpdateTag(target Tag, upd Tag) (Tag, error) {
	tx := GetDB().Begin()
	if err := tx.Error; err != nil {
		return upd, err
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	targetRec := tx.Model(&upd).Where(&target)
	if err := targetRec.Update(upd).Error; err != nil {
		tx.Rollback()
		return upd, err
	}
	if upd.ParentID < 1 {
		if err := targetRec.Update("parentId", gorm.Expr("NULL")).Error; err != nil {
			tx.Rollback()
			return upd, err
		}
	}
	return upd, tx.Commit().Error
}

func DeleteTag(target Tag, del Tag) (Tag, error) {
	err := GetDB().Where(&target).Delete(&del).Error
	return del, err
}
