package main

import (
	"books/src/controller"
	"books/src/middleware"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	engine := gin.Default()
	engine.Use(middleware.RecordUaAndTime)
	bookEngine := engine.Group("/book")
	{
		v1 := bookEngine.Group("v1")
		{
			v1.GET("/list", controller.BookList)
			v1.POST("/add", controller.BookAdd)
			v1.PUT("/update", controller.BookUpdate)
			v1.DELETE("/delete", controller.BookDelete)
		}
	}
	engine.Run(":3000")
}
