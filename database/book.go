package database

import (
	"RESTSAMPLE/models"

	"github.com/jinzhu/gorm"
)

func GetBooks(db *gorm.DB) ([]models.Book, error) {

	books := []models.Book{}
	query := db.Select("books.*")
	err := query.Find(&books).Error
	if err != nil {
		return nil, err
	}

	return books, nil

}
func GetBookbyID(db *gorm.DB, id string) error {
	return nil
}

func UpdateBook(db *gorm.DB, book models.Book) error {
	return nil
}
