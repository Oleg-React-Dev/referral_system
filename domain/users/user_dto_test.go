package users

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateInvalidEmail(t *testing.T) {
	user := User{Email: "  "}

	err := user.Validate()
	assert.EqualValues(t, err.Status, http.StatusBadRequest)
	assert.EqualValues(t, err.Message, "invalid email address")
	assert.EqualValues(t, err.Error, "bad_request")

	user.Email = "test#mail.ru"
	assert.EqualValues(t, err.Status, http.StatusBadRequest)
	assert.EqualValues(t, err.Message, "invalid email address")
	assert.EqualValues(t, err.Error, "bad_request")
}
func TestValidateInvalidPassword(t *testing.T) {
	user := User{Email: "test@mail.ru", Password: "  "}

	err := user.Validate()
	assert.EqualValues(t, err.Status, http.StatusBadRequest)
	assert.EqualValues(t, err.Message, "invalid password")
	assert.EqualValues(t, err.Error, "bad_request")
}
func TestValidateValidCredentials(t *testing.T) {
	user := User{Email: "test@mail.ru", Password: "123abc"}
	err := user.Validate()
	assert.Nil(t, err)
}
