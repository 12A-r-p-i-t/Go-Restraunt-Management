package helper

import (
	"log"
	"os"
	"time"

	"github.com/12A-r-p-i-t/restraunt-management/model"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

type Claims struct {
	Email      string `json:"email"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Password   string `json:"password"`
	Avatar     string `json:"avatar"`
	Phone      string `json:"phone"`
	jwt.RegisteredClaims
}

func GenerateToken(user model.User) (string, time.Time, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error in loading the .env file in helper.go")
	}
	secretKey := os.Getenv("SECRET_KEY")
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &Claims{
		Email:      user.Email,
		First_name: user.First_name,
		Last_name:  user.Last_name,
		Password:   user.Password,
		Avatar:     user.Avatar,
		Phone:      user.Phone,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		log.Fatal("Error in signing the token in helper.go")
		return "", expirationTime, err
	}
	return tokenString, expirationTime, nil
}
