package request

import (
	"github.com/brcodingdev/go-crud-users/internal/ports/response"
	"github.com/brcodingdev/go-crud-users/internal/util"
)

// UpdateUserRequest update user request
type UpdateUserRequest struct {
	Name      string `json:"name" example:"Jonh Doe"`
	BirthDate string `json:"birth_date" example:"1988-07-07"`
	Email     string `json:"email" example:"clebersonh@yahoo.com.br"`
	Address   string `json:"address" example:"135 W 45th St, New York, NY 10036, United States"`
	UserName  string `json:"username" example:"cleberson"`
	Password  string `json:"password" example:"123"`
} // @name UpdateUserRequest

// IsValid checks if request is valid
func (u UpdateUserRequest) IsValid() (bool, []response.Detail) {
	errors := make([]response.Detail, 0)

	if util.StringIsEmpty(u.Name) {
		errors = append(errors, response.MakeFieldDetail(
			"name",
			"name is empty",
		))
	}

	if util.StringIsEmpty(u.Email) {
		errors = append(errors, response.MakeFieldDetail(
			"email",
			"email is empty",
		))
	} else if !util.IsEmailValid(u.Email) {
		errors = append(errors, response.MakeFieldDetail(
			"email",
			"email is not valid",
		))
	}

	if util.StringIsEmpty(u.Address) {
		errors = append(errors, response.MakeFieldDetail(
			"address",
			"address is empty",
		))
	}

	if util.StringIsEmpty(u.UserName) {
		errors = append(errors, response.MakeFieldDetail(
			"username",
			"username is empty",
		))
	}

	if util.StringIsEmpty(u.Password) {
		errors = append(errors, response.MakeFieldDetail(
			"password",
			"password is empty",
		))
	}
	_, err := util.ParseDate(u.BirthDate)
	if err != nil {
		errors = append(errors, response.MakeFieldDetail(
			"birth_date",
			"birth_date is not valid",
		))
	}

	return len(errors) == 0, errors
}
