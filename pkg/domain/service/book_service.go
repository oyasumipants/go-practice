package service

import (
	"github.com/HazeyamaLab/go_crud/pkg/domain/model"
	"github.com/HazeyamaLab/go_crud/pkg/repository"
)

type BookService interface{
	Create(book model.Book) error
	FindAll() ([]model.Book, error)
	Update(book model.Book) error
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

func (b *bookService) Create(book model.Book) error{
	return b.BookRepository.Create(book)
}

func (b *bookService) Update(book model.Book) error{
	return b.BookRepository.Update(book)
}
