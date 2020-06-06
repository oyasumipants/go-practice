package server

import (
	"github.com/HazeyamaLab/go-crud/conf"
	"github.com/HazeyamaLab/go-crud/pkg/controller"
	"github.com/HazeyamaLab/go-crud/pkg/domain/service"
	"github.com/HazeyamaLab/go-crud/pkg/repository"
	"github.com/gin-gonic/gin"
)

func Init() {
	r := router()

	r.Run()
}

func router() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("view/*.html")
	bookController := bookInjector()
	r.GET("/", func(context *gin.Context) {
		context.HTML(200, "top.html", nil)
	})
	r.GET("/books", bookController.Index )
	r.GET("/books/:id/edit", bookController.Edit)
	r.POST("/books/:id/update",bookController.Update)
	r.POST("/book/create", bookController.Create)
	return r
}

func bookInjector() controller.BookController{
	db := conf.GetDB()
	bookRepository := repository.NewBookRepository(db)
	bookService := service.NewBookService(bookRepository)
	bookController := controller.NewBookController(bookService)
	return bookController
}


