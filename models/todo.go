package models

import (
	"todo/dao"
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func CreateTodo(todo *Todo) error {
	return dao.DB.Create(todo).Error
}

func GetAllTodo(todolist *[]Todo) error {
	return dao.DB.Find(todolist).Error
}

func GetTodo(id string, todo *Todo) error {
	return dao.DB.Where("id = ?", id).First(todo).Error
}

func UpdateTodo(todo *Todo) error {
	return dao.DB.Save(todo).Error
}

func DeleteTodo(id string) error {
	return dao.DB.Where("id = ?", id).Delete(Todo{}).Error
}
