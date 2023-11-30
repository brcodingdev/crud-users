package api

import (
	"errors"
	"github.com/brcodingdev/go-crud-users/internal/core/domain"
	aerrors "github.com/brcodingdev/go-crud-users/internal/errors"
	"github.com/brcodingdev/go-crud-users/internal/ports/request"
	_ "github.com/brcodingdev/go-crud-users/internal/ports/response" // used to swagger documentation
	"github.com/brcodingdev/go-crud-users/internal/util"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
)

// @Summary Update existing user
// @Description Update an existing user with the provided information
// @Tags users
// @Accept json
// @Produce json
// @Param user-id path string true "User ID" example(aa0e9b96-5f85-4383-a522-d6144602cb3e)
// @Param user body request.UpdateUserRequest true "User Data"
// @Success 204 "No Content"
// @Failure 400 {object} response.ErrorResponse "bad request"
// @Failure 404 {object} response.ErrorResponse "user not found"
// @Failure 409 {object} response.ErrorResponse "username or email already exists"
// @Failure 500 {object} response.ErrorResponse "internal server error"
// @Router /users/{user-id} [put]
func (api UserAPI) handleUpdate(
	w http.ResponseWriter,
	r *http.Request,
) {
	vars := mux.Vars(r)
	userID := vars["user-id"]
	if !util.IsUUID(userID) {
		makeErrResponse(
			"invalid uuid",
			http.StatusBadRequest,
			w,
		)
		return
	}

	updateUserRequest := request.UpdateUserRequest{}
	err := parseBody(r, &updateUserRequest)
	if err != nil {
		makeErrResponse(
			"could not parse user",
			http.StatusBadRequest,
			w,
		)
		return
	}
	isValid, details := updateUserRequest.IsValid()
	if !isValid {
		makeErrResponseWithDetails(
			details,
			http.StatusBadRequest,
			w,
		)
		return
	}
	// uuid already validated
	userUUID, _ := uuid.Parse(userID)
	userDomain := buildUserDomainFromUpdateUser(updateUserRequest, userUUID)
	err = api.service.UpdateUser(userDomain)
	if err != nil {
		var notFoundError *aerrors.RecordNotFoundError
		if ok := errors.As(err, &notFoundError); ok {
			makeErrResponse(
				"user not found",
				http.StatusNotFound,
				w,
			)
			return
		}

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
			"could not update user",
			http.StatusInternalServerError,
			w,
		)
		return
	}
	noContent(w)
}

// INTERNAL FUNCTIONS

// build user domain from update user request
func buildUserDomainFromUpdateUser(
	updateUserRequest request.UpdateUserRequest,
	userID uuid.UUID,
) domain.User {
	//birthDate already validated
	birthDate, _ := util.ParseDate(updateUserRequest.BirthDate)
	return domain.User{
		ID:        &userID,
		Name:      updateUserRequest.Name,
		BirthDate: birthDate,
		Email:     updateUserRequest.Email,
		Address:   updateUserRequest.Address,
		UserName:  updateUserRequest.UserName,
		Password:  updateUserRequest.Password,
	}
}
