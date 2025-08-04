package routers

import (
	"todo/controller"

	"github.com/gin-gonic/gin"
)

func SetRouter() *gin.Engine {
	r := gin.Default()
	//加载静态文件
	r.Static("/static", "static")
	//加载模板
	r.LoadHTMLGlob("templates/*")
	r.GET("/", controller.IndexHandler)

	v1Group := r.Group("v1")
	{
		//增
		v1Group.POST("/todo", controller.CreateTodo)
		//查
		v1Group.GET("/todo", controller.GetTodo)
		v1Group.GET("/todo/:id", func(c *gin.Context) {

		})
		//改
		v1Group.PUT("/todo/:id", controller.GetTodo)
		//删
		v1Group.DELETE("/todo/:id", controller.DeleteTodo)
	}
	return r
}
