package models

type LabModel struct {
	ID       uint     `gorm:"column:id;primary_key;auto_increment"`
	LabName  string   `gorm:"column:lab_name;unique;index:idx_lab_name;not null;<-:create"`
	LabPass  string   `gorm:"column:lab_pass;not null;<-:create"`
	UserList []string `gorm:"column:user_list;type:json"`
}

func (LabModel) TableName() string {
	return "lab"
}
