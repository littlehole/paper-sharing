package models

type LabModel struct {
	ID       uint     `gorm:"column:id;primary_key;auto_increment"`
	LabName  string   `gorm:"column:lab_name;unique"`
	UserList []string `gorm:"column:user_list;type:json"`
}

func (LabModel) TableName() string {
	return "lab"
}
