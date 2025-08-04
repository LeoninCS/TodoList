package controller

import (
	"net/http"
	"todo/models"

	"github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func CreateTodo(c *gin.Context) {
	//取数据
	var todo models.Todo
	c.BindJSON(&todo)
	//存入数据库
	if err := models.CreateTodo(&todo); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func GetTodo(c *gin.Context) {
	var todolist []models.Todo
	if err := models.GetAllTodo(&todolist); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todolist)
	}
}

func UpdateTodo(c *gin.Context) {
	//获取id
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	//数据库查询
	var todo models.Todo
	if err := models.GetTodo(id, &todo); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}
	//绑定前端传过来的数据
	c.BindJSON(&todo)
	//数据库中更改数据
	if err := models.UpdateTodo(&todo); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func DeleteTodo(c *gin.Context) {
	//获取id
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	//数据库查询
	if err := models.DeleteTodo(id); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{id: "deleted"})
	}
}
