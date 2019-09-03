package internal

import(
	"errors"
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

func GetTagParents(child Tag) ([]Tag, error) {
	var checkedTags []Tag
	checkedTags = append(checkedTags, child)
	if child.ParentID < 1 {
		return checkedTags, nil
	}
	var allTags []Tag
	err := GetDB().Where(&Tag{ UserID: child.UserID }).Find(&allTags).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return checkedTags, nil
		} else {
			return checkedTags, err
		}
	}
	for i := 0; i < len(checkedTags); i++ {
		nextId := checkedTags[len(checkedTags) - 1].ParentID
		if nextId < 1 {
			return checkedTags, nil
		}
		for _, v := range allTags {
			if nextId == v.ID {
				for _, w := range checkedTags {
					if v.ID == w.ID {
						return checkedTags, errors.New("タグの親子関係が循環しています")
					}
				}
				checkedTags = append(checkedTags, v)
				break
			}
		}
	}
	return checkedTags, errors.New("指定された親タグが見つかりませんでした")
}

func GetTagChildren(parent Tag) ([]Tag, error) {
	var children []Tag
	children = append(children, parent)
	var queue []int
	queue = append(queue, parent.ID)
	var allTags []Tag
	if err := GetDB().Where(&Tag{ UserID: parent.UserID }).Find(&allTags).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return children, nil
		} else {
			return children, err
		}
	}
	for {
		appendId := queue[0]
		queue = queue[1:]
		for _, v := range allTags {
			if (appendId == v.ParentID) {
				children = append(children, v)
				queue = append(queue, v.ID)
			}
		}
		if len(queue) == 0 {
			break
		}
	}
	return children, err
}
