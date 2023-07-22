package models

import (
	"diary_api/database"
	"errors"

	"gorm.io/gorm"
)

type Entry struct {
	gorm.Model
	Content string `gorm:"type:text" json:"content"`
	UserID  uint
}

func (entry *Entry) Save() (*Entry, error) {
	result := database.Database.Create(&entry)
	if result.RowsAffected == 0 {
		return &Entry{}, errors.New("fail add entry to database")
	}

	if result.Error != nil {
		return &Entry{}, result.Error
	}

	return entry, nil
}
