package models

import (
	"gorm.io/gorm"
	"strconv"
	"time"
)

type Todos struct {
	TodoId          int64     `gorm:"primaryKey" json:"id"`
	ActivityGroupId int64     `json:"activity_group_id"`
	Title           string    `json:"title"`
	Priority        string    `json:"priority"`
	IsActive        bool      `json:"is_active"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}

func (Todos) TableName() string {
	return "todos"
}

func (todos *Todos) Create(db *gorm.DB) map[string]interface{} {
	todos.Priority = "very-high"
	err := db.Create(&todos).Error
	if err != nil {
		return Message("Failed", err.Error())
	}

	resp := Message("Success", "Success")
	resp["data"] = todos

	return resp
}

func (todos *Todos) Update(id int64, db *gorm.DB) map[string]interface{} {
	vtodos := &Todos{}
	err := db.Where("todo_id = ?", id).First(&vtodos).Error
	if err != nil {
		return Message("Not Found", "Todo with ID"+strconv.FormatInt(id, 10)+"Not Found")
	}

	vtodos.Title = todos.Title
	vtodos.Priority = todos.Priority
	vtodos.ActivityGroupId = todos.ActivityGroupId
	vtodos.IsActive = todos.IsActive
	err = db.Save(&vtodos).Error
	if err != nil {
		return Message("Failed", err.Error())
	}

	resp := Message("Success", "Success")
	resp["data"] = vtodos

	return resp
}

func GetTodos(id int64, db *gorm.DB) map[string]interface{} {
	todos := &Todos{}
	err := db.Where("todo_id = ?", id).First(&todos).Error
	if err != nil {
		return Message("Not Found", "Todo with ID"+strconv.FormatInt(id, 10)+"Not Found")
	}
	resp := Message("Success", "Success")
	resp["data"] = todos

	return resp
}

func GetAllTodos(activity_group_id int64, db *gorm.DB) map[string]interface{} {
	todos := make([]*Todos, 0)
	db.Where("activity_group_id = ?", activity_group_id).Find(&todos)
	resp := Message("Success", "Success")

	resp["data"] = todos

	return resp
}

func DeleteTodos(id int64, db *gorm.DB) map[string]interface{} {
	todos := &Todos{}
	err := db.Where("todo_id = ?", id).First(&todos).Error
	if err != nil {
		return Message("Not Found", "Todo with ID"+strconv.FormatInt(id, 10)+"Not Found")
	}

	err = db.Where("id = ?", id).Delete(&todos).Error
	if err != nil {
		return Message("Failed", err.Error())
	}

	resp := Message("Success", "Success")
	return resp
}
