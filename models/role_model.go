package models

type Role struct {
	ID   uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name string `gorm:"type:varchar(255);not null" json:"name"`
}
