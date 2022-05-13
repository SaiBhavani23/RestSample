package models

import "github.com/jinzhu/gorm"

type Book struct {
	gorm.Model
	author    string
	name      string
	pageCount string
}
