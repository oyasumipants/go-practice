package repository

import (
	"github.com/HazeyamaLab/go_crud/pkg/domain/model"
	"github.com/jinzhu/gorm"
)

type BookRepository interface{
	Create(book model.Book) error
	FindAll() ([]model.Book, error)
	Update(book model.Book) error
}


type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository{
	return &bookRepository{db: db}
}

//CRUDのR READで情報を取得する この関数は全ての書籍のデータを取る
func (b *bookRepository) FindAll() ([]model.Book, error){
	var books []model.Book
	err := b.db.Find(&books).Error
	return books, err
}

func (b *bookRepository) Create(book model.Book) error{
	return b.db.Create(&book).Error
}

func (b *bookRepository) Update(book model.Book) error {
	return b.db.Save(&book).Error
}