package api

import (
	"encoding/json"
	"errors"
	"github.com/brcodingdev/go-crud-users/internal/core/domain"
	aerrors "github.com/brcodingdev/go-crud-users/internal/errors"
	"github.com/brcodingdev/go-crud-users/internal/ports/request"
	"github.com/brcodingdev/go-crud-users/internal/ports/response"
	"github.com/brcodingdev/go-crud-users/internal/util"
	"net/http"
)

// @Summary Create new user
// @Description Creates a new user with the provided information
// @Tags users
// @Accept json
// @Produce json
// @Param user body request.CreateUserRequest true "User Data"
// @Success 200 {object} response.CreateUserResponse
// @Failure 400 {object} response.ErrorResponse "bad request"
// @Failure 409 {object} response.ErrorResponse "username or email already exists"
// @Failure 500 {object} response.ErrorResponse "internal server error"
// @Router /users [post]
func (api UserAPI) handleCreate(
	w http.ResponseWriter,
	r *http.Request,
) {
	createUserRequest := request.CreateUserRequest{}
	err := parseBody(r, &createUserRequest)
	if err != nil {
		makeErrResponse(
			"could not parse request body",
			http.StatusBadRequest,
			w,
		)
		return
	}
	isValid, details := createUserRequest.IsValid()
	if !isValid {
		makeErrResponseWithDetails(
			details,
			http.StatusBadRequest,
			w,
		)
		return
	}
	userDomain := buildUserDomainFromCreateUser(createUserRequest)
	createdUser, err := api.service.CreateUser(userDomain)
	if err != nil {
		var dupError *aerrors.DuplicateEntryError
		if ok := errors.As(err, &dupError); ok {
			makeErrResponse(
				"username or email already exists",
				http.StatusConflict,
				w,
			)
			return
		}

		makeErrResponse(
			"could not create user",
			http.StatusInternalServerError,
			w,
		)
		return
	}

	res := buildCreateUserResponse(createdUser)
	data, err := json.Marshal(res)
	if err != nil {
		makeErrResponse(
			"could not marshal response",
			http.StatusInternalServerError,
			w,
		)
		return
	}

	ok(data, w)
}

// INTERNAL FUNCTIONS

// build user domain from create user request
func buildUserDomainFromCreateUser(
	createUserRequest request.CreateUserRequest,
) domain.User {
	//birthDate already validated
	birthDate, _ := util.ParseDate(createUserRequest.BirthDate)
	return domain.User{
		Name:      createUserRequest.Name,
		BirthDate: birthDate,
		Email:     createUserRequest.Email,
		Address:   createUserRequest.Address,
		UserName:  createUserRequest.UserName,
		Password:  createUserRequest.Password,
	}
}

// build user response from user domain
func buildCreateUserResponse(
	userDomain *domain.User,
) response.CreateUserResponse {
	return response.CreateUserResponse{
		ID:        *userDomain.ID,
		Name:      userDomain.Name,
		BirthDate: userDomain.BirthDate,
		Email:     userDomain.Email,
		Address:   userDomain.Address,
		UserName:  userDomain.UserName,
		CreatedAt: *userDomain.CreatedAt,
		UpdatedAt: *userDomain.CreatedAt,
	}
}
