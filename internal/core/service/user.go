package service

import (
	"github.com/brcodingdev/go-crud-users/internal/adapters/model"
	"github.com/brcodingdev/go-crud-users/internal/adapters/repository"
	"github.com/brcodingdev/go-crud-users/internal/core/domain"
	"github.com/brcodingdev/go-crud-users/internal/util"
	"github.com/google/uuid"
)

// User ...
type User interface {
	// CreateUser creates user
	CreateUser(user domain.User) (*domain.User, error)
	// GetUserByID gets user by its ID
	GetUserByID(ID string) (*domain.User, error)
	// UpdateUser updates the user
	UpdateUser(user domain.User) error
	// DeleteUser deletes the user
	DeleteUser(ID string) error
}

// UserService ...
type UserService struct {
	userRepository repository.User
}

// NewUserService ...
func NewUserService(
	userRepository repository.User,
) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

// CreateUser ...
func (service UserService) CreateUser(
	user domain.User,
) (*domain.User, error) {
	hashPass, err := util.GenerateHashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	userModel := buildUserModel(user, uuid.New(), string(hashPass))
	userModelCreated, err := service.userRepository.Create(&userModel)
	if err != nil {
		return nil, err
	}

	user.ID = &userModelCreated.ID
	user.CreatedAt = &userModelCreated.CreatedAt
	user.UpdatedAt = &userModelCreated.UpdatedAt

	return &user, nil
}

// GetUserByID ...
func (service UserService) GetUserByID(
	ID string,
) (*domain.User, error) {
	userModel, err := service.userRepository.GetByID(ID)
	if err != nil {
		return nil, err
	}
	userDomain := buildUserDomain(*userModel)
	return &userDomain, nil
}

// UpdateUser ...
func (service UserService) UpdateUser(user domain.User) error {
	userModel, err := service.userRepository.GetByID(user.ID.String())
	if err != nil {
		return err
	}

	hashPass, err := util.GenerateHashPassword(user.Password)
	if err != nil {
		return err
	}
	userModel.ID = *user.ID
	userModel.Name = user.Name
	userModel.BirthDate = user.BirthDate
	userModel.Email = user.Email
	userModel.Address = user.Address
	userModel.UserName = user.UserName
	userModel.Password = string(hashPass)

	err = service.userRepository.Update(userModel)
	if err != nil {
		return err
	}
	return nil
}

// DeleteUser ...
func (service UserService) DeleteUser(ID string) error {
	_, err := service.userRepository.GetByID(ID)
	if err != nil {
		return err
	}
	return service.userRepository.Delete(ID)
}

// INTERNAL FUNCTIONS
func buildUserModel(
	user domain.User,
	ID uuid.UUID,
	password string) model.User {
	return model.User{
		ID:        ID,
		Name:      user.Name,
		BirthDate: user.BirthDate,
		Email:     user.Email,
		Address:   user.Address,
		UserName:  user.UserName,
		Password:  password,
	}
}

func buildUserDomain(user model.User) domain.User {
	return domain.User{
		ID:        &user.ID,
		Name:      user.Name,
		BirthDate: user.BirthDate,
		Email:     user.Email,
		Address:   user.Address,
		UserName:  user.UserName,
		CreatedAt: &user.CreatedAt,
		UpdatedAt: &user.UpdatedAt,
	}
}
