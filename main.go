package main

import (
	"fmt"
	"todo/dao"
	"todo/models"
	"todo/routers"
)

func main() {

	//连接数据库

	err := dao.InitMysql()
	if err != nil {
		fmt.Println("wrong")
		panic(err)
	}
	dao.DB.AutoMigrate(&models.Todo{})
	defer dao.Close()

	r := routers.SetRouter()

	r.Run()
}
