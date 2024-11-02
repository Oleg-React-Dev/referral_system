package referral_codes

import (
	"fmt"
	"referral_app/datasources/postegresql/users_db"
	"referral_app/domain/users"
	"referral_app/logger"
	"referral_app/utils/errors"
	"strings"
)

const (
	queryInsertReferralCode     = "INSERT INTO referral_codes (user_id, expires_at) VALUES ($1, $2) RETURNING code;"
	queryDeleteReferralCode     = "DELETE FROM referral_codes WHERE user_id=$1 RETURNING code;"
	queryGetReferralCodeByEmail = "SELECT code, expires_at FROM referral_codes join users using(user_id) WHERE email=$1;"
	queryFinedById              = "SELECT user_id, email, created_at, updated_at FROM users JOIN referrals ON user_id = referral_id WHERE referrer_id = $1"

	duplicateErr = "duplicate key value violates unique constraint"
	notFoundErr  = "no rows in result set"
)

func (rc *ReferralCode) Save() *errors.RestErr {
	saveErr := users_db.Db.QueryRow(queryInsertReferralCode, rc.UserId, rc.ExpiresAt).Scan(&rc.Code)

	if saveErr != nil {
		if strings.Contains(saveErr.Error(), duplicateErr) {
			return errors.NewBadRequestError("the code already exists for given user")
		}
		logger.Error("error when trying to save referral code", saveErr)
		return errors.NewInternalServerError("database error")
	}
	return nil
}

func (rc *ReferralCode) Delete() *errors.RestErr {
	deleteErr := users_db.Db.QueryRow(queryDeleteReferralCode, rc.UserId).Scan(&rc.Code)
	if deleteErr != nil {
		if strings.Contains(deleteErr.Error(), notFoundErr) {
			return errors.NewBadRequestError("there is no referral code for given user")
		}
		logger.Error("error when trying to delete referral code", deleteErr)
		return errors.NewInternalServerError("database error")
	}
	return nil
}

func (rc *ReferralCode) GetCodeByEmail(email string) *errors.RestErr {
	err := users_db.Db.QueryRow(queryGetReferralCodeByEmail, email).Scan(&rc.Code, &rc.ExpiresAt)
	if err != nil {
		if strings.Contains(err.Error(), notFoundErr) {
			return errors.NewBadRequestError("there is no referral code for given email address")
		}
		logger.Error("error when trying to get referral code", err)
		return errors.NewInternalServerError("database error")
	}
	return nil
}

func (u *ReferralCode) FindById(id string) (users.Users, *errors.RestErr) {
	rows, err := users_db.Db.Query(queryFinedById, id)
	if err != nil {
		logger.Error("error when trying to find referrals by referrer id", err)
		return nil, errors.NewInternalServerError("database error")
	}

	results := make(users.Users, 0)
	for rows.Next() {
		var user users.User
		if err := rows.Scan(&user.UserId, &user.Email, &user.CreatedAt, &user.UpdatedAt); err != nil {
			logger.Error("error when trying to scan referral row into user struct", err)
			return nil, errors.NewInternalServerError("database error")
		}
		results = append(results, user)
	}

	if len(results) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no users matching id %s", id))
	}
	return results, nil
}
