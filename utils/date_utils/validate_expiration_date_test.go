package date_utils

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateExpirationDateConstants(t *testing.T) {
	assert.EqualValues(t, apiDbLayout, "2006-01-02 15:04:05.999999+00", "should match the layout")
}

func TestValidateExpirationDateEmptyString(t *testing.T) {
	err := ValidateExpirationDate("  ")
	assert.EqualValues(t, err.Status, http.StatusBadRequest)
	assert.EqualValues(t, err.Message, "invalid expiration date: cannot be empty")
	assert.EqualValues(t, err.Error, "bad_request")
}
func TestValidateExpirationDateInvalidFormat(t *testing.T) {
	err := ValidateExpirationDate("2006-01-02 15:04:05")
	assert.EqualValues(t, err.Status, http.StatusBadRequest)
	assert.EqualValues(t, err.Message, "invalid expiration date: must be in 'YYYY-MM-DD HH:MM:SS.mmmmmm+00' format")
	assert.EqualValues(t, err.Error, "bad_request")
}
func TestValidateExpirationDateExpired(t *testing.T) {
	err := ValidateExpirationDate("2000-01-02 15:04:05.999999+00")
	assert.EqualValues(t, err.Status, http.StatusBadRequest)
	assert.EqualValues(t, err.Message, "invalid expiration date: must be in the future")
	assert.EqualValues(t, err.Error, "bad_request")
}
func TestValidateExpirationDateOk(t *testing.T) {
	err := ValidateExpirationDate("2030-01-02 15:04:05.999999+00")
	assert.Nil(t, err)
}
