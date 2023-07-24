package controller

import (
	"ListOfItemsProject/initDB"
	"ListOfItemsProject/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// TodoGet
// 获取数据库内的文件
func TodoGet(c *gin.Context) {
	var todolist []*models.Todo

	err := initDB.Db.Find(&todolist).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, todolist)
	}
}

// TodoUpdate
// 对某一项数据进行更新
func TodoUpdate(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"status": "same id",
		})
	}
	var todo models.Todo
	err := initDB.Db.Where("id = ?", id).First(&todo).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = c.BindJSON(&todo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	}

	err = initDB.Db.Save(&todo).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

// TodoDelete
// 删除某项数据
func TodoDelete(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"status": "can't find id",
		})
	}
	var todo models.Todo
	err := initDB.Db.Where("id = ?", id).First(&todo).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = initDB.Db.Where("id = ?", id).Delete(&models.Todo{}).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{id: "deleted"})
	}
}

// TodoPost
// 新增数据
func TodoPost(c *gin.Context) {
	var todo models.Todo
	err := c.BindJSON(&todo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = initDB.Db.Create(&todo).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, todo)
	}
}
