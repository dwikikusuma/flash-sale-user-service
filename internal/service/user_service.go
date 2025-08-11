package service

import (
	"errors"
	"user-management-service/internal/entity"
	"user-management-service/internal/repository"
)

type UserService interface {
	GetUserDetails(userID int64) (*entity.User, error)
	CreateUser(user *entity.User) (*entity.User, error)
	Login(request *entity.LoginRequest) (*entity.User, error)
}

type userService struct {
	UserRepository repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		UserRepository: userRepo,
	}
}

func (u *userService) GetUserDetails(userID int64) (*entity.User, error) {
	// This is a placeholder implementation.
	// In a real application, you would fetch the user details from a database or another source.
	userDetail, err := u.UserRepository.GetUserByID(userID)
	if err != nil {
		return nil, err
	}
	if userDetail == nil {
		return nil, errors.New("user not found")
	}

	return userDetail, nil
}

func (u *userService) CreateUser(user *entity.User) (*entity.User, error) {
	// This is a placeholder implementation.
	// In a real application, you would save the user to a database or another source.
	userDetail, err := u.UserRepository.CreateUser(user)
	if err != nil {
		return nil, err
	}

	if userDetail == nil {
		return nil, errors.New("failed to create user")
	}
	return userDetail, nil
}

func (u *userService) Login(request *entity.LoginRequest) (*entity.User, error) {
	// This is a placeholder implementation.
	// In a real application, you would validate the user's credentials against a database or another source.
	userDetail, err := u.UserRepository.GetUserByEmail(request.Username)
	if err != nil {
		return nil, err
	}

	if userDetail == nil || userDetail.Password != request.Password {
		return nil, errors.New("invalid username or password")
	}
	return userDetail, nil
}
