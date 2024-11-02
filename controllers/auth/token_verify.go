package auth

import (
	"fmt"
	"os"
	"referral_app/utils/errors"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Claims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	jwt.StandardClaims
}

func TokenVerifyMiddleWare(next gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		authToken, authErr := getToken(c.GetHeader("Authorization"))
		if authErr != nil {
			c.JSON(authErr.Status, authErr)
			return
		}

		claims, err := parseToken(authToken)
		if err != nil {
			restErr := errors.NewUnauthorizedError(err.Error())
			c.JSON(restErr.Status, restErr)
			return
		}

		if claims.UserID != "" && claims.Email != "" {
			c.Set("user_id", claims.UserID)
			c.Set("email", claims.Email)
			next(c)
		} else {
			restErr := errors.NewUnauthorizedError("Invalid token.")
			c.JSON(restErr.Status, restErr)
			return
		}
	}
}

func getToken(authHeader string) (string, *errors.RestErr) {
	bearerToken := strings.Split(authHeader, " ")
	if len(bearerToken) == 2 {
		return bearerToken[1], nil

	}
	return "", errors.NewUnauthorizedError("Invalid token.")
}

func parseToken(authToken string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(authToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		if claims.ExpiresAt < time.Now().Unix() {
			return nil, fmt.Errorf("token expired")
		}
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}
