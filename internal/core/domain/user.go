package domain

import (
	"github.com/google/uuid"
	"time"
)

// User ...
type User struct {
	ID        *uuid.UUID
	Name      string
	BirthDate time.Time
	Email     string
	Address   string
	UserName  string
	Password  string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
