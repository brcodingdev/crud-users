package api

import (
	"encoding/json"
	"github.com/brcodingdev/go-crud-users/internal/ports/response"
	"io"
	"net/http"
)

func parseBody(r *http.Request, x interface{}) error {
	if body, err := io.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal(body, x); err != nil {
			return err
		}
	}
	return nil
}

func ok(res []byte, httpCode int, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)
	w.Write(res)
}

func noContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

func sendErrResponse(errorResponse response.ErrorResponse, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(errorResponse.Code)
	data, _ := json.Marshal(errorResponse)
	w.Write(data)
}

func makeErrResponse(detail string, httpCode int, w http.ResponseWriter) {
	sendErrResponse(response.ErrorResponse{
		Errors: []response.Detail{
			response.MakeDetail(detail),
		},
		Code: httpCode,
	}, w)
}

func makeErrResponseWithDetails(
	errors []response.Detail,
	httpCode int,
	w http.ResponseWriter,
) {
	sendErrResponse(response.ErrorResponse{
		Errors: errors,
		Code:   httpCode,
	}, w)
}
