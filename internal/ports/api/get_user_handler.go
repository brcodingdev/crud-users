package api

import (
	"encoding/json"
	"errors"
	"github.com/brcodingdev/go-crud-users/internal/core/domain"
	aerrors "github.com/brcodingdev/go-crud-users/internal/errors"
	"github.com/brcodingdev/go-crud-users/internal/ports/response"
	"github.com/brcodingdev/go-crud-users/internal/util"
	"github.com/gorilla/mux"
	"net/http"
)

// @Summary Get user
// @Description Get an existing user by their ID
// @Tags users
// @Accept json
// @Produce json
// @Param user-id path string true "User ID" example(aa0e9b96-5f85-4383-a522-d6144602cb3e)
// @Success 200 {object} response.UserResponse
// @Failure 400 {object} response.ErrorResponse "bad request"
// @Failure 404 {object} response.ErrorResponse "user not found"
// @Failure 500 {object} response.ErrorResponse "internal server error"
// @Router /users/{user-id} [get]
func (api UserAPI) handleGet(
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

	user, err := api.service.GetUserByID(userID)
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

		makeErrResponse(
			"could not find user",
			http.StatusInternalServerError,
			w,
		)
		return
	}

	res := buildUserResponse(user)

	data, err := json.Marshal(res)
	if err != nil {
		makeErrResponse(
			"could not marshal user",
			http.StatusInternalServerError,
			w,
		)
		return
	}

	ok(data, http.StatusOK, w)
}

// INTERNAL FUNCTIONS

// build user response from user domain
func buildUserResponse(userDomain *domain.User) response.UserResponse {
	return response.UserResponse{
		ID:        *userDomain.ID,
		Name:      userDomain.Name,
		Age:       util.CalculateAge(userDomain.BirthDate),
		Email:     userDomain.Email,
		Address:   userDomain.Address,
		UserName:  userDomain.UserName,
		CreatedAt: *userDomain.CreatedAt,
		UpdatedAt: *userDomain.CreatedAt,
	}
}
