package models

import (
	"gorm.io/gorm"
	"strconv"
	"time"
)

type Activities struct {
	ActivityId int64     `gorm:"primaryKey" json:"id"`
	Title      string    `json:"title"`
	Email      string    `json:"email"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

func (Activities) TableName() string {
	return "activities"
}

func (activities *Activities) Create(db *gorm.DB) map[string]interface{} {
	vactivities := activities
	query := db.Create(&vactivities)
	if query.Error != nil {
		return Message("Failed", query.Error.Error())
	}

	resp := Message("Success", "Success")
	resp["data"] = vactivities

	return resp
}

func (activities *Activities) Update(id int64, db *gorm.DB) map[string]interface{} {
	vactivities := &Activities{}
	err := db.Where("activity_id = ?", id).First(&vactivities).Error
	if err != nil {
		return Message("Not Found", "Activity with ID "+strconv.FormatInt(id, 10)+" Not Found")
	}

	vactivities.Title = activities.Title
	err = db.Save(&vactivities).Error
	if err != nil {
		return Message("Failed", err.Error())
	}

	resp := Message("Success", "Success")
	resp["data"] = vactivities

	return resp
}

func GetActivities(id int64, db *gorm.DB) map[string]interface{} {
	activities := &Activities{}
	err := db.Where("activity_id = ?", id).First(&activities).Error
	if err != nil {
		return Message("Not Found", "Activity with ID"+strconv.FormatInt(id, 10)+"Not Found")
	}
	resp := Message("Success", "Success")
	resp["data"] = activities

	return resp
}

func GetAllActivities(db *gorm.DB) map[string]interface{} {
	activities := make([]*Activities, 0)
	db.Find(&activities)
	resp := Message("Success", "Success")

	resp["data"] = activities

	return resp
}

func DeleteActivities(id int64, db *gorm.DB) map[string]interface{} {
	activities := &Activities{}
	err := db.Where("activity_id = ?", id).First(&activities).Error
	if err != nil {
		return Message("Not Found", "Activity with ID"+strconv.FormatInt(id, 10)+"Not Found")
	}

	err = db.Where("id = ?", id).Delete(&activities).Error
	if err != nil {
		return Message("Failed", err.Error())
	}

	resp := Message("Success", "Success")
	return resp
}
