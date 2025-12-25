package response

import (
	"encoding/json"
	"net/http"
)

type (
	Error       string
	ContentType string
)

const (
	// Error types
	ErrInvalidInput = Error("INVALID_INPUT")
	ErrNotFound     = Error("NOT_FOUND")
	ErrInternal     = Error("INTERNAL_SERVER_ERROR")

	// Success types
	SuccessOk        = "OK"
	SuccessCreated   = "CREATED"
	SuccessUpdated   = "UPDATED"
	SuccessNoContent = "NO_CONTENT"
	SuccessDeleted   = "DELETED"

	// Content Types
	ContentTypeJSON = ContentType("application/json")
)

var (
	Errors = map[string]struct {
		Code    int
		Message string
		Status  string
	}{
		"INVALID_INPUT": {
			Code:    400,
			Message: "The input provided is invalid.",
			Status:  string(ErrInvalidInput),
		},
		"NOT_FOUND": {
			Code:    404,
			Message: "The requested resource was not found.",
			Status:  string(ErrNotFound),
		},
		"INTERNAL_SERVER_ERROR": {
			Code:    500,
			Message: "An internal server error occurred.",
			Status:  string(ErrInternal),
		},
	}

	SuccessMessages = map[string]struct {
		Code    int
		Message string
		Status  string
	}{
		"OK": {
			Code:    200,
			Message: "Request was successful.",
			Status:  string(SuccessOk),
		},
		"CREATED": {
			Code:    201,
			Message: "Resource was created successfully.",
			Status:  string(SuccessCreated),
		},
		"UPDATED": {
			Code:    200,
			Message: "Resource was updated successfully.",
			Status:  string(SuccessOk),
		},
		"NO_CONTENT": {
			Code:    204,
			Message: "Request was successful but there is no content to return.",
			Status:  string(SuccessNoContent),
		},
		"DELETED": {
			Code:    200,
			Message: "Resource was deleted successfully.",
			Status:  string(SuccessDeleted),
		},
	}
) // var

func JSON(w http.ResponseWriter, status int, payload any) {
	out, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", string(ContentTypeJSON))
	w.WriteHeader(status)
	w.Write(out)
}

func Err(w http.ResponseWriter, err Error, message ...string) {
	out, _ := json.Marshal(Errors[string(err)])
	w.Header().Set("Content-Type", string(ContentTypeJSON))
	w.WriteHeader(Errors[string(err)].Code)
	w.Write(out)
}

func Success(w http.ResponseWriter, success string, message ...string) {
	out, _ := json.Marshal(SuccessMessages[string(success)])
	w.Header().Set("Content-Type", string(ContentTypeJSON))
	w.WriteHeader(SuccessMessages[string(success)].Code)
	w.Write(out)
}
