package services

import (
	"clean-architecture-api/lib"
	"clean-architecture-api/models"
	"clean-architecture-api/repository"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// UserService service layer
type UserService struct {
	logger          lib.Logger
	repository      repository.UserRepository
	paginationScope *gorm.DB
}

// NewUserService creates a new userservice
func NewUserService(
	logger lib.Logger,
	userRepository repository.UserRepository,
) *UserService {
	return &UserService{
		logger:     logger,
		repository: userRepository,
	}
}

// WithTrx delegates transaction to repository database
func (s UserService) WithTrx(trxHandle *gorm.DB) UserService {
	s.repository = s.repository.WithTrx(trxHandle)
	return s
}

// PaginationScope
func (s UserService) SetPaginationScope(scope func(*gorm.DB) *gorm.DB) UserService {
	s.paginationScope = s.repository.WithTrx(s.repository.Scopes(scope)).DB
	return s
}

// GetOneUser gets one user
func (s UserService) GetOneUser(userID lib.BinaryUUID) (user models.User, err error) {
	return user, s.repository.First(&user, "id = ?", userID).Error
}

// GetUserByEmail gets one user by email
func (s UserService) GetUserByEmail(email string) (user models.User, err error) {
	return user, s.repository.First(&user, "Email = ?", email).Error
}

// GetAllUser get all the user
func (s UserService) GetAllUser() (response map[string]interface{}, err error) {
	var users []models.User
	var count int64

	err = s.repository.WithTrx(s.paginationScope).Find(&users).Offset(-1).Limit(-1).Count(&count).Error
	if err != nil {
		return nil, err
	}

	return gin.H{"data": users, "count": count}, nil
}

// UpdateUser updates the user
func (s UserService) UpdateUser(user *models.User) error {
	return s.repository.Save(&user).Error
}

// DeleteUser deletes the user
func (s UserService) DeleteUser(uuid lib.BinaryUUID) error {
	return s.repository.Where("id = ?", uuid).Delete(&models.User{}).Error
}

// Create creates a new user
func (s UserService) Create(user *models.User) error {
	return s.repository.Create(&user).Error
}

func (s UserService) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (s UserService) CompareHashAndPassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		s.logger.Error(err)
		return false
	}
	return true
}

func (s UserService) GenerateJWTToken(user models.User) (string, error) {
	// Generate a JWT token
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": user.ID,
		"email":  user.Email,
		"iat":    time.Now().Unix(),
		"exp":    time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	return tokenString, err
}

func (s UserService) ValidateJWTToken(tokenString string) (jwt.MapClaims, error) {
	// decode and validate
	// Parse takes the token string and a function for looking up the key.
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("failed to decode token")
}
