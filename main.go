package main

import (
	"ListOfItemsProject/controller"
	"ListOfItemsProject/initDB"
	"ListOfItemsProject/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {

	initDB.InitDB()
	err := initDB.Db.AutoMigrate(&models.Todo{})
	if err != nil {
		log.Fatal(err)
	}

	route := gin.Default()
	route.Static("/static", "static")
	route.LoadHTMLGlob("index.html")
	route.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	routeGroup := route.Group("v1")
	{
		//查找
		routeGroup.GET("/todo", controller.TodoGet)
		//修改
		routeGroup.PUT("/todo/:id", controller.TodoUpdate)
		//删除
		routeGroup.DELETE("/todo/:id", controller.TodoDelete)
		//添加
		routeGroup.POST("/todo", controller.TodoPost)
	}

	err = route.Run()
	if err != nil {
		log.Fatal(err)
	}
}
