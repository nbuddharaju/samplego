package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	userId   string `json:"userid" gorm:"text;not null;default:null`
	userName string `json:"username" gorm:"text;not null;default:null`
}
