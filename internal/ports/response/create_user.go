package response

import (
	"github.com/google/uuid"
	"time"
)

// CreateUserResponse create user response
type CreateUserResponse struct {
	ID        uuid.UUID `json:"id" example:"29dda872-8f1d-11ee-8e7f-66c94c3ada1d"`
	Name      string    `json:"name" example:"Jonh Doe"`
	BirthDate time.Time `json:"birth_date" example:"1988-07-07"`
	Email     string    `json:"email" example:"clebersonh@yahoo.com.br"`
	Address   string    `json:"address" example:"Jo135 W 45th St, New York, NY 10036, United States"`
	UserName  string    `json:"username" example:"cleberson"`
	CreatedAt time.Time `json:"created_at" example:"2023-11-29T20:47:18.923-03:00"`
	UpdatedAt time.Time `json:"updated_at" example:"2023-11-29T20:47:18.923-03:00"`
} // @name CreateUserResponse
