package request

import (
	"github.com/brcodingdev/go-crud-users/internal/ports/response"
	"github.com/brcodingdev/go-crud-users/internal/util"
)

// CreateUserRequest request to create user
type CreateUserRequest struct {
	Name      string `json:"name" example:"Jonh Doe"`
	BirthDate string `json:"birth_date" example:"1988-07-07"`
	Email     string `json:"email" example:"clebersonh@yahoo.com.br"`
	Address   string `json:"address" example:"135 W 45th St, New York, NY 10036, United States"`
	UserName  string `json:"username" example:"cleberson"`
	Password  string `json:"password" example:"123"`
} // @name CreateUserRequest

// IsValid checks if request is valid
func (c CreateUserRequest) IsValid() (bool, []response.Detail) {
	errors := make([]response.Detail, 0)

	if util.StringIsEmpty(c.Name) {
		errors = append(errors, response.MakeFieldDetail(
			"name",
			"name is empty",
		))
	}

	if util.StringIsEmpty(c.Email) {
		errors = append(errors, response.MakeFieldDetail(
			"email",
			"email is empty",
		))
	} else if !util.IsEmailValid(c.Email) {
		errors = append(errors, response.MakeFieldDetail(
			"email",
			"email is not valid",
		))
	}

	if util.StringIsEmpty(c.Address) {
		errors = append(errors, response.MakeFieldDetail(
			"address",
			"address is empty",
		))
	}

	if util.StringIsEmpty(c.UserName) {
		errors = append(errors, response.MakeFieldDetail(
			"username",
			"username is empty",
		))
	}

	if util.StringIsEmpty(c.Password) {
		errors = append(errors, response.MakeFieldDetail(
			"password",
			"password is empty",
		))
	}
	_, err := util.ParseDate(c.BirthDate)
	if err != nil {
		errors = append(errors, response.MakeFieldDetail(
			"birth_date",
			"birth_date is not valid",
		))
	}

	return len(errors) == 0, errors
}
