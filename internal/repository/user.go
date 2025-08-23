package repository

import (
	"context"
	"errors"
	"user-management-service/infrasturcture/log"
	"user-management-service/internal/entity"

	"gorm.io/gorm"
)

// UserRepository defines the interface for user-related data operations.
type UserRepository interface {
	// GetUserByID retrieves a user by their ID.
	GetUserByID(ctx context.Context, id int64) (*entity.User, error)
	// CreateUser adds a new user to the repository.
	CreateUser(ctx context.Context, user *entity.User) error
	// UpdateUser modifies an existing user in the repository.
	UpdateUser(ctx context.Context, user *entity.User) (*entity.User, error)
	// DeleteUser removes a user from the repository by their ID.
	DeleteUser(ctx context.Context, id int64) error
	// GetUserByEmail retrieves a user by their email address.
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
}

// userRepository is a concrete implementation of the UserRepository interface.
type userRepository struct {
	db *gorm.DB
}

// NewUserRepository creates and returns a new instance of userRepository.
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

// GetUserByID retrieves a user by their ID from the in-memory data store.
// Returns the user if found, or nil if the user does not exist.
func (r *userRepository) GetUserByID(ctx context.Context, id int64) (*entity.User, error) {
	var user entity.User
	err := r.db.Table("users").WithContext(ctx).Where("id = ?", id).First(&user).Error
	if err != nil {
		log.Logger.Error().Err(err).Msg("Failed to get user by ID")
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil

}

// CreateUser adds a new user to the in-memory data store.
// Assigns a new ID to the user and returns the created user.
func (r *userRepository) CreateUser(ctx context.Context, user *entity.User) error {
	return r.db.Table("users").WithContext(ctx).Create(user).Error
}

// UpdateUser updates an existing user in the in-memory data store.
// Returns the updated user.
func (r *userRepository) UpdateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	err := r.db.Table("users").WithContext(ctx).Save(user).Error
	if err != nil {
		log.Logger.Error().Err(err).Msg("Failed to update user")
		return nil, err
	}

	return user, nil
}

// DeleteUser removes a user from the in-memory data store by their ID.
// Returns nil if the operation is successful.
func (r *userRepository) DeleteUser(ctx context.Context, id int64) error {
	isExists, err := r.GetUserByID(ctx, id)
	if err != nil {
		log.Logger.Error().Err(err).Msg("Failed to check user existence before deletion")
		return err
	}

	if isExists == nil {
		log.Logger.Warn().Int64("userID", id).Msg("User not found for deletion")
		return gorm.ErrRecordNotFound
	}

	return r.db.Table("users").WithContext(ctx).Where("id = ?", id).Delete(&entity.User{}).Error
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user entity.User
	err := r.db.Table("users").WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		log.Logger.Error().Err(err).Msg("Failed to get user by email")
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
