package model

type Test struct {
	ID   int    `gorm:"primaryKey" json:"id""`
	Name string `json:"name"`
}
