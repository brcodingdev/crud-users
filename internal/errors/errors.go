package errors

// DuplicateEntryError ...
type DuplicateEntryError struct {
	message string
}

func (e *DuplicateEntryError) Error() string {
	return e.message
}

// NewDuplicateEntryError ...
func NewDuplicateEntryError(msg string) error {
	return &DuplicateEntryError{message: msg}
}

// RecordNotFoundError ...
type RecordNotFoundError struct {
	message string
}

func (e *RecordNotFoundError) Error() string {
	return e.message
}

// NewRecordNotFoundError ...
func NewRecordNotFoundError(msg string) error {
	return &RecordNotFoundError{message: msg}
}
