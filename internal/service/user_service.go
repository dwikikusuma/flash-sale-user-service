package service

import (
	"context"
	"errors"
	"time"
	"user-management-service/config"
	"user-management-service/infrasturcture/log"
	"user-management-service/internal/entity"
	"user-management-service/internal/repository"
	iUtils "user-management-service/utils"

	"github.com/golang-jwt/jwt/v5"
)

type UserService interface {
	GetUserDetails(ctx context.Context, userID int64) (*entity.User, error)
	CreateUser(ctx context.Context, user *entity.User) error
	Login(ctx context.Context, request *entity.LoginRequest) (string, error)
}

type userService struct {
	UserRepository repository.UserRepository
	AppConfig      config.Config
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		UserRepository: userRepo,
	}
}

func (u *userService) GetUserDetails(ctx context.Context, userID int64) (*entity.User, error) {
	// This is a placeholder implementation.
	// In a real application, you would fetch the user details from a database or another source.
	userDetail, err := u.UserRepository.GetUserByID(ctx, userID)
	if err != nil {
		log.Logger.Error().Err(err).Msg("Failed to get user details")
		return nil, err
	}
	if userDetail == nil {
		log.Logger.Warn().Int64("userID", userID).Msg("User not found")
		return nil, errors.New("user not found")
	}

	return userDetail, nil
}

func (u *userService) CreateUser(ctx context.Context, user *entity.User) error {
	// This is a placeholder implementation.
	// In a real application, you would save the user to a database or another source.
	hashedPassword, err := iUtils.GenerateHashedPassword(user.Password)
	if err != nil {
		log.Logger.Error().Err(err).Msg("Failed to hash password")
		return err
	}

	user.Password = hashedPassword
	err = u.UserRepository.CreateUser(ctx, user)
	if err != nil {
		log.Logger.Error().Err(err).Msg("Failed to create user")
		return err
	}
	return nil
}

func (u *userService) Login(ctx context.Context, request *entity.LoginRequest) (string, error) {
	// This is a placeholder implementation.
	// In a real application, you would validate the user's credentials against a database or another source.
	userDetail, err := u.UserRepository.GetUserByEmail(ctx, request.Username)
	if err != nil {
		log.Logger.Error().Err(err).Msg("Failed to get user by email")
		return "nil", err
	}

	if userDetail == nil || userDetail.Password != request.Password {
		log.Logger.Warn().Str("username", request.Username).Msg("Invalid username or password")
		return "", errors.New("invalid username or password")
	}

	isPasswordMatch, err := iUtils.ComparedPassword(request.Password, userDetail.Password)
	if err != nil {
		log.Logger.Error().Err(err).Msg("Failed to compare password")
		return "", err
	}

	if !isPasswordMatch {
		log.Logger.Warn().Str("username", request.Username).Msg("Invalid username or password")
		return "", errors.New("invalid username or password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  userDetail.ID,
		"username": userDetail.Username,
		"email":    userDetail.Email,
		"exp":      jwt.NewNumericDate(time.Now().Add(2 * time.Hour)),
	})

	signedToken, err := token.SignedString([]byte(u.AppConfig.Secret.JWTSecret))
	if err != nil {
		log.Logger.Error().Err(err).Msg("Failed to generate token")
		return "", err
	}

	return signedToken, nil
}
