// services/login_service.go

package service

import (
	"auth/domain/models"
	"auth/domain/repositories"
	"auth/dto/request"
	"errors"
	"fmt"
	"github.com/azerpost/dashboard-lib/errs"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

type AuthService interface {
	Login(request.LoginRequest) (string, error)
	SignUp(request.SignUpRequest) (*models.User, *errs.AppError)
}

type authService struct {
	UserRepository repositories.UserRepository
	Secret         *[]byte
}

func NewLoginService(userRepository repositories.UserRepository) *authService {
	return &authService{
		UserRepository: userRepository,
	}
}

func (s *authService) Login(req request.LoginRequest) (string, error) {
	user, err := s.UserRepository.FindByEmail(req.Email)
	if err != nil {
		return "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return "", errors.New("Invalid email or password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.ID,
		"username": user.Email,
		"exp":      time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", errors.New("Failed to create token")
	}

	return tokenString, nil
}

func (s *authService) SignUp(request request.SignUpRequest) (*models.User, *errs.AppError) {
	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)
	if err != nil {
		return nil, errs.NewInternalServerError("Failed to hash password")
	}

	user := models.User{Email: request.Email, Username: request.Username, Password: string(hash)}

	if user.IsMailValid() {
		fmt.Println("is valid")
	}

	createdUser, err := s.UserRepository.CreateUser(&user)
	if err != nil {
		return nil, errs.NewInternalServerError("Failed to create user")
	}

	return createdUser, nil
}

//func getUserFromToken(tokenString string) (*jwt.Token, error) {
//	token, err := jwt.ParseWithClaims(tokenString, &domain.AccessTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
//		return []byte(os.Getenv("SECRET")), nil
//	})
//	if err != nil {
//		logger.Error("Error while parsing token: " + err.Error())
//		return nil, err
//	}
//	return token, nil
//}
