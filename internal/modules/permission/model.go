package permission

type Permission struct {
	ID   uint   `json:"id" gorm:"primaryKey;autoIncrement;column:id_permission"`
	Key  string `json:"key" gorm:"type:varchar(100);unique;not null"`
	Desc string `json:"description" gorm:"type:text"`
}

func (Permission) TableName() string { return "permissions" }
