package models

type Usr struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func (Usr) TableName() string {
	return "usr"
}
