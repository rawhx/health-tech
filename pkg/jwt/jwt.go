package jwt

import (
	"errors"
	"health-tech/models"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Interface interface {
	CreateJWTToken(userID string) (string, error)
	ValidateToken(tokenString string) (string, error)
	GetLoginUser(c *gin.Context) (*models.User, error)
}

type jsonWebToken struct {
	SecretKey   string
	ExpiredTime time.Duration
}

type Claims struct {
	UserID  string
	jwt.RegisteredClaims
}

func Init() Interface {
	secretKey := os.Getenv("JWT_KEY")
	expiredTime, err := strconv.Atoi(os.Getenv("JWT_EXP_TIME"))
	if err != nil {
		log.Fatalf("error init jwt %v", err)
	}

	return &jsonWebToken{
		SecretKey:   secretKey,
		ExpiredTime: time.Duration(expiredTime) * time.Hour,
	}
}

func (j *jsonWebToken) CreateJWTToken(userID string) (string, error) {
	claims := &Claims{
		UserID:  userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.ExpiredTime)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(j.SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (j *jsonWebToken) ValidateToken(tokenString string) (string, error) {
	var (
		claim  Claims
		userID string
	)

	token, err := jwt.ParseWithClaims(tokenString, &claim, func(t *jwt.Token) (interface{}, error) {
		return []byte(j.SecretKey), nil
	})

	if err != nil {
		return userID, err
	}

	if !token.Valid {
		return userID, errors.New("token is not valid")
	}

	userID = claim.UserID
	return userID, nil
}

func (j *jsonWebToken) GetLoginUser(c *gin.Context) (*models.User, error) {
	user, ok := c.Get("user")
	if !ok {
		return &models.User{}, errors.New("failed to get user login")
	}

	return user.(*models.User), nil
}