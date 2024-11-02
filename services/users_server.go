package services

import (
	"os"
	"referral_app/domain/user_jwt"
	"referral_app/domain/users"
	"referral_app/logger"
	"referral_app/utils/errors"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

const (
	expirationTime = 24
)

var UserService usersServiceInterface = &userService{}

type userService struct{}

type usersServiceInterface interface {
	CreateUser(users.RegisterUserRequest) (*users.User, *errors.RestErr)
	LoginUser(users.LoginUserRequest) (*user_jwt.JWT, *errors.RestErr)
}

func (s *userService) CreateUser(u users.RegisterUserRequest) (*users.User, *errors.RestErr) {
	user := users.User{Email: u.Email, Password: u.Password, ReferralCode: u.ReferralCode}

	if err := user.Validate(); err != nil {
		return nil, err
	}

	hash, errCrypt := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	if errCrypt != nil {
		logger.Error("error when hashing password", errCrypt)
		return nil, errors.NewInternalServerError("internal server error")
	}

	user.Password = string(hash)

	if err := user.Save(); err != nil {
		return nil, err
	}

	user.ReferralCode = strings.TrimSpace(user.ReferralCode)
	if user.ReferralCode != "" {
		user.SaveReferral()
	}
	return &user, nil
}

func (s *userService) LoginUser(request users.LoginUserRequest) (*user_jwt.JWT, *errors.RestErr) {

	dao := &users.User{
		Email:    request.Email,
		Password: request.Password,
	}

	if err := dao.Validate(); err != nil {
		return nil, err
	}

	password := request.Password

	if err := dao.FinedByEmail(); err != nil {
		return nil, err
	}

	hashedPassword := dao.Password
	passwordErr := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if passwordErr != nil {
		return nil, errors.NewNotFoundError("invalid user credentials")
	}

	token, tokenErr := GenerateToken(dao)
	if tokenErr != nil {
		return nil, errors.NewInternalServerError("database error")

	}
	return &user_jwt.JWT{Token: token}, nil
}

func GenerateToken(user *users.User) (string, *errors.RestErr) {
	secret := os.Getenv("SECRET")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":   user.Email,
		"user_id": user.UserId,
		"iss":     "referral_system",
		"exp":     time.Now().Add(expirationTime * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		logger.Error("error when trying generate token", err)
		return "", errors.NewInternalServerError("internal server error")
	}
	return tokenString, nil
}
