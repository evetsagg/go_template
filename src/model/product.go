package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// todo should separate json and db entities
type Product struct {
	Id           string    `json:"Id" gorm:"primary_key"` //if integer, use this to autoincrement `gorm:"primaryKey;auto_increment;not null"` uint
	ProductTitle string    `json:"ProductTitle"`
	ProductDesc  string    `json:"ProductDesc"`
	Update       time.Time `gorm:"autoCreateTime"`
	Create       time.Time `gorm:"autoUpdateTime"`
	Increment    int       `gorm:"auto_increment,not null"`
	//gorm:index for indexing
}

func (product *Product) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	product.Id = uuid.NewString()
	return
}
