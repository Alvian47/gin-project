package models

type Book struct {
	ID     uint   `json:"id" gorm:"primaryKey, autoIncrement"`
	Title  string `json:"title"`
	Author string `json:"author"`
}
