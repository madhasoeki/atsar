package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID            uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	NamaLengkap   string         `gorm:"type:varchar(255);not null" json:"nama_lengkap"`
	NamaPanggilan string         `gorm:"type:varchar(255);not null" json:"nama_panggilan"`
	Email         string         `gorm:"type:varchar(255);not null;unique" json:"email"`
	Password      string         `gorm:"type:varchar(255);not null" json:"-"`
	RoleID        uint           `gorm:"not null" json:"role_id"`
	Role          Role           `gorm:"foreignKey:RoleID" json:"role"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
}
