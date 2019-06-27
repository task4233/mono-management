package internal

type ReqSearch struct {
	Name  string
	TagID int `json:"tagId" gorm:"tagId"`
}
