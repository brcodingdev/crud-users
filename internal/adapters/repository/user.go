package repository

import (
	"errors"
	"github.com/brcodingdev/go-crud-users/internal/adapters/model"
	aerrors "github.com/brcodingdev/go-crud-users/internal/errors"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

// User interface to handle DB user operations
type User interface {
	// Create creates user
	Create(user *model.User) (*model.User, error)
	// GetByID gets user by ID
	GetByID(ID string) (*model.User, error)
	// Update updates the user
	Update(user *model.User) error
	// Delete deletes the user
	Delete(ID string) error
}

// UserRepository ...
type userRepository struct {
	db *gorm.DB
}

// NewUserUserRepository create a repository instance
func NewUserUserRepository(db *gorm.DB) User {
	return &userRepository{
		db: db,
	}
}

// Create ...
func (u *userRepository) Create(user *model.User) (*model.User, error) {
	db := u.db.Create(user)
	if err := db.Error; err != nil {
		if isDuplicateEntryError(err) {
			return nil, aerrors.NewDuplicateEntryError(
				err.Error(),
			)
		}

		return nil, err
	}

	return user, nil
}

// GetByID ...
func (u *userRepository) GetByID(ID string) (*model.User, error) {
	var user *model.User
	db := u.db.First(&user, "ID", ID)
	if err := db.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, aerrors.NewRecordNotFoundError(
				err.Error(),
			)
		}
		return nil, err
	}
	return user, nil
}

// Update ...
func (u *userRepository) Update(user *model.User) error {
	if err := db.Save(user).Error; err != nil {
		if isDuplicateEntryError(err) {
			return aerrors.NewDuplicateEntryError(
				err.Error(),
			)
		}
		return err
	}
	return nil
}

// Delete ...
func (u *userRepository) Delete(ID string) error {
	var user *model.User
	// Unscoped is a hard delete, if you want soft then remove it
	if err := db.Unscoped().Delete(&user, "ID", ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return aerrors.NewRecordNotFoundError(
				err.Error(),
			)
		}
		return err
	}
	return nil
}

// INTERNAL FUNCTIONS

// checks if there is a constraint error
func isDuplicateEntryError(err error) bool {
	var mysqlErr *mysql.MySQLError
	ok := errors.As(err, &mysqlErr)
	if !ok {
		return false
	}
	return mysqlErr.Number == 1062
}
