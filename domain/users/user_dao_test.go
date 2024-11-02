package users

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserDaoConstants(t *testing.T) {
	assert.EqualValues(t, queryInsertUser, "INSERT INTO users (email, password) VALUES ($1, $2) RETURNING user_id, created_at, updated_at;", "should match the query string")
	assert.EqualValues(t, queryGetUserPassword, "SELECT password, user_id FROM users WHERE email=$1", "should match the query string")
	assert.EqualValues(t, queryAddReferral, "SELECT add_referral($1, $2);", "should match the query string")
}
