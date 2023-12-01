package api

import (
	"errors"
	aerrors "github.com/brcodingdev/go-crud-users/internal/errors"
	_ "github.com/brcodingdev/go-crud-users/internal/ports/response" // used to swagger documentation
	"github.com/brcodingdev/go-crud-users/internal/util"
	"github.com/gorilla/mux"
	"net/http"
)

// @Summary Delete existing user
// @Description Delete existing user
// @Tags users
// @Accept json
// @Produce json
// @Param user-id path string true "User ID" example(aa0e9b96-5f85-4383-a522-d6144602cb3e)
// @Success 204 "No Content"
// @Failure 400 {object} response.ErrorResponse "bad request"
// @Failure 404 {object} response.ErrorResponse "user not found"
// @Failure 500 {object} response.ErrorResponse "internal server error"
// @Router /users/{user-id} [delete]
func (api *UserAPI) handleDelete(
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

	// uuid already validated
	err := api.service.DeleteUser(userID)
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
			"could not delete user",
			http.StatusInternalServerError,
			w,
		)
		return
	}
	noContent(w)
}
