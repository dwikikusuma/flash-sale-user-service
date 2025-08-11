package repository

import "user-management-service/internal/entity"

// UserRepository defines the interface for user-related data operations.
type UserRepository interface {
	// GetUserByID retrieves a user by their ID.
	GetUserByID(id int64) (*entity.User, error)
	// CreateUser adds a new user to the repository.
	CreateUser(user *entity.User) (*entity.User, error)
	// UpdateUser modifies an existing user in the repository.
	UpdateUser(user *entity.User) (*entity.User, error)
	// DeleteUser removes a user from the repository by their ID.
	DeleteUser(id int64) error
	// GetUserByEmail retrieves a user by their email address.
	GetUserByEmail(email string) (*entity.User, error)
}

// userRepository is a concrete implementation of the UserRepository interface.
type userRepository struct {
}

// NewUserRepository creates and returns a new instance of userRepository.
func NewUserRepository() UserRepository {
	return &userRepository{}
}

// users is an in-memory data store simulating a database for user entities.
var users = map[int64]entity.User{
	1: {
		ID:       1,
		Username: "john_doe",
		Email:    "jhon_doe@example.com",
		Password: "hashed_password",
	},
	2: {
		ID:       2,
		Username: "jane_doe",
		Email:    "jane_doe@example.com",
		Password: "hashed_password",
	},
}

// GetUserByID retrieves a user by their ID from the in-memory data store.
// Returns the user if found, or nil if the user does not exist.
func (r *userRepository) GetUserByID(id int64) (*entity.User, error) {
	user, ok := users[id]
	if !ok {
		return nil, nil // User not found
	}
	return &user, nil
}

// CreateUser adds a new user to the in-memory data store.
// Assigns a new ID to the user and returns the created user.
func (r *userRepository) CreateUser(user *entity.User) (*entity.User, error) {
	user.ID = int64(len(users) + 1) // Simulating an auto-generated ID
	users[user.ID] = *user
	return user, nil
}

// UpdateUser updates an existing user in the in-memory data store.
// Returns the updated user.
func (r *userRepository) UpdateUser(user *entity.User) (*entity.User, error) {
	users[user.ID] = *user
	return user, nil
}

// DeleteUser removes a user from the in-memory data store by their ID.
// Returns nil if the operation is successful.
func (r *userRepository) DeleteUser(id int64) error {
	delete(users, id)
	return nil
}

func (r *userRepository) GetUserByEmail(email string) (*entity.User, error) {
	for _, user := range users {
		if user.Email == email {
			return &user, nil
		}
	}
	return nil, nil
}
