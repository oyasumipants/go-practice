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
	r.GET("/", bookController.Index )
	r.POST("/create", bookController.Create)
	return r
}

func bookInjector() controller.BookController{
	db := conf.GetDB()
	repo := repository.NewBookRepository(db)
	ser := service.NewBookService(repo)
	ctr := controller.NewBookController(ser)
	return ctr
}


