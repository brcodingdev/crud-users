package response

// ErrorResponse error handling response
type ErrorResponse struct {
	// Errors an slice of returned errors
	Errors []Detail `json:"errors"`
	Code   int      `json:"code"`
} // @name ErrorResponse

// Detail defines the struct to return errors
type Detail struct {
	Field  string `json:"field,omitempty"`
	Detail string `json:"detail"`
} // @name Detail

// MakeDetail creates a new detail error
func MakeDetail(detail string) Detail {
	return MakeFieldDetail("", detail)
}

// MakeFieldDetail creates a new detail error with field and detail
// attributes filled in.
func MakeFieldDetail(field string, detail string) Detail {
	if len(field) > 0 {
		return Detail{Field: field, Detail: detail}
	}

	return Detail{
		Detail: detail,
	}
}
