package users

import (
	"database/sql"
	"referral_app/datasources/postegresql/users_db"
	"referral_app/logger"
	"referral_app/utils/errors"
)

const (
	queryInsertUser      = "INSERT INTO users (email, password) VALUES ($1, $2) RETURNING user_id, created_at, updated_at;"
	queryGetUserPassword = "SELECT password, user_id FROM users WHERE email=$1"
	queryAddReferral     = "SELECT add_referral($1, $2);"
)

func (u *User) Save() *errors.RestErr {
	saveErr := users_db.Db.QueryRow(queryInsertUser, u.Email, u.Password).Scan(&u.UserId, &u.CreatedAt, &u.UpdatedAt)

	if saveErr != nil {
		logger.Error("error when trying to save user", saveErr)
		return errors.NewInternalServerError("database error")
	}
	u.Password = ""
	return nil
}

func (u *User) FinedByEmail() *errors.RestErr {
	row := users_db.Db.QueryRow(queryGetUserPassword, u.Email)
	scanErr := row.Scan(&u.Password, &u.UserId)
	if scanErr != nil {
		if scanErr == sql.ErrNoRows {
			return errors.NewNotFoundError("invalid user credentials")
		}

		logger.Error("error when trying to get user by email", scanErr)
		return errors.NewInternalServerError("database error")
	}
	return nil
}

func (u *User) SaveReferral() *errors.RestErr {
	var success bool
	err := users_db.Db.QueryRow(queryAddReferral, u.UserId, u.ReferralCode).Scan(&success)

	if err != nil {
		logger.Error("error when trying to add referral", err)
		return errors.NewInternalServerError("database error")
	}
	if !success {
		logger.Info("error when trying to add referral, Referral code is wrong or expired")
		u.ReferralCode = "Referral code is wrong or expired"
	}
	u.Password = ""
	return nil
}
