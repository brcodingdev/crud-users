package api

import (
	"github.com/brcodingdev/go-crud-users/internal/core/service"
	"github.com/gorilla/mux"
	"net/http"
)

// UserAPI ...
type UserAPI struct {
	router  *mux.Router
	service service.User
}

// NewUserAPI creates a user api
func NewUserAPI(
	router *mux.Router,
	service service.User,
) *UserAPI {
	return &UserAPI{
		router:  router,
		service: service,
	}
}

// RegisterRoutes register user routes pointing to /users
func (api *UserAPI) RegisterRoutes() {
	subRouter := api.router.PathPrefix("/users").Subrouter()

	subRouter.HandleFunc(
		"",
		api.handleCreate,
	).Methods(http.MethodPost)

	subRouter.HandleFunc(
		"/{user-id}",
		api.handleGet,
	).Methods(http.MethodGet)

	subRouter.HandleFunc(
		"/{user-id}",
		api.handleUpdate,
	).Methods(http.MethodPut)

	subRouter.HandleFunc(
		"/{user-id}",
		api.handleDelete,
	).Methods(http.MethodDelete)
}
