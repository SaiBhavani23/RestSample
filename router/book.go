package router

import (
	"RESTSAMPLE/database"
	"RESTSAMPLE/handler"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	api := handler.Handler{
		DB: database.GetDB(),
	}
	r.GET("/books", api.GetBooksHndlr)
	return nil
}
