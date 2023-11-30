package api

import (
	_ "github.com/brcodingdev/go-crud-users/docs" // used for swagger docs
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

// RegisterRoute ...
func RegisterRoute() *mux.Router {
	r := mux.NewRouter()
	// Swagger endpoint
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	return r
}
