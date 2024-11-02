package errors

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBadRequestError(t *testing.T) {
	err := NewBadRequestError("bad request error massage")
	assert.EqualValues(t, err.Status, http.StatusBadRequest)
	assert.EqualValues(t, err.Message, "bad request error massage")
	assert.EqualValues(t, err.Error, "bad_request")
}
func TestNewNotFoundError(t *testing.T) {
	err := NewNotFoundError("not found error massage")
	assert.EqualValues(t, err.Status, http.StatusNotFound)
	assert.EqualValues(t, err.Message, "not found error massage")
	assert.EqualValues(t, err.Error, "not_found")
}
func TestNewInternalServerError(t *testing.T) {
	err := NewInternalServerError("internal server error massage")
	assert.EqualValues(t, err.Status, http.StatusInternalServerError)
	assert.EqualValues(t, err.Message, "internal server error massage")
	assert.EqualValues(t, err.Error, "internal_server_error")
}
func TestNewUnauthorizedError(t *testing.T) {
	err := NewUnauthorizedError("unauthorized error massage")
	assert.EqualValues(t, err.Status, http.StatusUnauthorized)
	assert.EqualValues(t, err.Message, "unauthorized error massage")
	assert.EqualValues(t, err.Error, "unauthorized_error")
}
