package permission

type Permission struct {
	ID     uint   `json:"id" gorm:"primaryKey;autoIncrement;column:id_permission"`
	Key    string `json:"key" gorm:"type:varchar(100);unique;not null"`
	Desc   string `json:"description" gorm:"type:text"`
	Status string `json:"status" gorm:"type:varchar(20);not null;default:'active'"`
}

func (Permission) TableName() string { return "permissions" }
