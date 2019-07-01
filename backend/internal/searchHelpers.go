package internal

type ReqSearch struct {
	Name  string `json:"name"`
	TagID int `json:"tagId" gorm:"column:tagId"`
}
