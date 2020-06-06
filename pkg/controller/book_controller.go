package controller

import (
	"github.com/HazeyamaLab/go_crud/pkg/domain/model"
	"github.com/HazeyamaLab/go_crud/pkg/domain/service"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

type BookController interface {
	Create(ctx *gin.Context)
	Index(ctx *gin.Context)
}

type bookController struct {
	bookService service.BookService
}

func NewBookController(b service.BookService) BookController{
	return &bookController{bookService: b}
}

func(b *bookController) Index(ctx *gin.Context){
	books, err := b.bookService.FindAll()
	if err != nil {
		ctx.HTML(500,"500.html",nil)
	}

	ctx.HTML(200, "index.html", gin.H{"books": books })
}

func(b *bookController) Create(ctx *gin.Context){
	title := ctx.PostForm("title")
	author := ctx.PostForm("author")

	//intに変換
	price, err := strconv.Atoi(ctx.PostForm("price"))
	if err != nil {
		panic(err)
	}
	book := model.Book{Title: title, Price: price, Author: author}

	err = b.bookService.Create(book)
	if err != nil {
		log.Println(err)
		ctx.HTML(500, "500.html", nil)
		return
	}

	ctx.Redirect(302, "/")
}


func(b *bookController) Update(ctx *gin.Context){
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		panic(err)
	}

	title := ctx.PostForm("title")
	author := ctx.PostForm("author")


	//intに変換
	price, err := strconv.Atoi(ctx.PostForm("price"))
	if err != nil {
		panic(err)
	}


	book := model.Book{ID: id, Title: title, Price: price, Author: author}

	err = b.bookService.Update(book)
	if err != nil {
		ctx.HTML(500, "500.html", nil)
		return
	}

	ctx.Redirect(302, "/")
}


