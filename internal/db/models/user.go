package models

import (
	"time"
)

type UserModel struct {
	ID        uint     `gorm:"column:id;primary_key;auto_increment"`
	LabName   string   `gorm:"column:lab_name;size:60;index:idx_lab_name_grade_name,priority:1"`
	Grade     string   `gorm:"column:grade;size:10;index:idx_lab_name_grade_name,priority:2"`
	Name      string   `gorm:"column:name;size:30;index:idx_lab_name_grade_name,priority:3"`
	Username  string   `gorm:"column:username;size:30;unique;index:idx_username"`
	Password  string   `gorm:"column:password;size:255;not null"`
	ShareList []string `gorm:"column:share_list;type:json"`
	PaperList []string `gorm:"column:paper_list;type:json"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (UserModel) TableName() string {
	return "users"
}
