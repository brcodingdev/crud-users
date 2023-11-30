package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// User ...
type User struct {
	ID        uuid.UUID `gorm:"type:char(36);primaryKey"`
	Name      string    `gorm:"not null;index:idx_user_name"`
	BirthDate time.Time `gorm:"type:date;not null"`
	Email     string    `gorm:"type:varchar(255);not null;uniqueIndex:idx_user_email"`
	Address   string    `gorm:"type:varchar(500);not null"`
	UserName  string    `gorm:"type:varchar(50);not null;uniqueIndex:idx_user_username"`
	Password  string    `gorm:"type:varchar(255);not null"`
	gorm.Model
}
