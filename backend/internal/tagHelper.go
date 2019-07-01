package internal


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

func UpdateTag(tag Tag) (Tag, error) {
	err := GetDB().Model(&tag).Update(tag).Error
	return tag, err
}

func DeleteTag(tag Tag) (Tag, error) {
	err := GetDB().Delete(&tag).Error
	return tag, err
}
