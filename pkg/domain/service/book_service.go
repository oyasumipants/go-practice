package service

import (
	"github.com/HazeyamaLab/go-crud/pkg/domain/model"
	"github.com/HazeyamaLab/go-crud/pkg/repository"
)

type BookService interface{
	Create(book model.Book) error
	FindAll() ([]model.Book, error)
	FindOne(id int) (model.Book, error)
	Update(book model.Book) error
	Delete(id int) error
}

type bookService struct {
	repository.BookRepository
}

func NewBookService(r repository.BookRepository) BookService {
	return &bookService{r}
}

func (b *bookService) FindAll() ([]model.Book, error){
	return b.BookRepository.FindAll()
}

func (b *bookService) FindOne(id int) (model.Book, error){
	return b.BookRepository.FindOne(id)
}

func (b *bookService) Create(book model.Book) error{
	return b.BookRepository.Create(book)
}

func (b *bookService) Update(book model.Book) error{
	return b.BookRepository.Update(book)
}

func (b *bookService) Delete(id int) error{
	return b.BookRepository.Delete(id)
}
