package repository_test

import (
	"errors"
	"github.com/brcodingdev/go-crud-users/internal/adapters/model"
	"github.com/brcodingdev/go-crud-users/internal/adapters/repository/mocks"
	"github.com/brcodingdev/go-crud-users/internal/core/domain"
	"github.com/brcodingdev/go-crud-users/internal/core/service"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestCreate(t *testing.T) {
	cases := map[string]struct {
		user          domain.User
		userExpected  model.User
		errorExpected error
	}{
		"success": {
			user: domain.User{
				Name:      "Jonh Doe",
				BirthDate: time.Date(1988, 07, 07, 0, 0, 0, 0, time.UTC),
				Email:     "test@test.com",
				Address:   "135 W 45th St, New York, NY 10036, United States",
				UserName:  "cleberson",
				Password:  "123",
			},
			userExpected: model.User{
				ID:        uuid.New(),
				Name:      "Jonh Doe",
				BirthDate: time.Date(1988, 07, 07, 0, 0, 0, 0, time.UTC),
				Email:     "test@test.com",
				Address:   "135 W 45th St, New York, NY 10036, United States",
				UserName:  "cleberson",
				Password:  "123",
			},
		},
	}

	for caseTitle, tc := range cases {
		t.Run(caseTitle, func(t *testing.T) {
			repository := mocks.User{}

			repository.On(
				"Create",
				mock.Anything,
			).Return(&tc.userExpected, tc.errorExpected)

			userService := service.NewUserService(&repository)

			userDomain, err := userService.CreateUser(tc.user)
			assert.NoError(t, err)
			assert.NotEmpty(t, userDomain.ID)
			assert.NotEmpty(t, userDomain.Name)
			assert.False(t, userDomain.BirthDate.IsZero())
			assert.NotEmpty(t, userDomain.Email)
			assert.NotEmpty(t, userDomain.Address)
			assert.NotEmpty(t, userDomain.UserName)
			assert.NotEmpty(t, userDomain.Password)
		})
	}
}

func TestCreateError(t *testing.T) {
	cases := map[string]struct {
		user          domain.User
		userExpected  model.User
		errorExpected error
	}{
		"error": {
			user: domain.User{
				Name:      "Jonh Doe",
				BirthDate: time.Date(1988, 07, 07, 0, 0, 0, 0, time.UTC),
				Email:     "test@test.com",
				Address:   "135 W 45th St, New York, NY 10036, United States",
				UserName:  "cleberson",
				Password:  "123",
			},
			errorExpected: errors.New("email or username already exists"),
		},
	}

	for caseTitle, tc := range cases {
		t.Run(caseTitle, func(t *testing.T) {
			repository := mocks.User{}

			repository.On(
				"Create",
				mock.Anything,
			).Return(nil, tc.errorExpected)

			userService := service.NewUserService(&repository)

			userDomain, err := userService.CreateUser(tc.user)
			assert.Error(t, err)
			assert.Nil(t, userDomain)
		})
	}
}

func TestUpdate(t *testing.T) {
	userID := uuid.New()
	cases := map[string]struct {
		user          domain.User
		userFound     model.User
		errorExpected error
	}{
		"success": {
			user: domain.User{
				ID:        &userID,
				Name:      "Jonh Doe",
				BirthDate: time.Date(1988, 07, 07, 0, 0, 0, 0, time.UTC),
				Email:     "test@test.com",
				Address:   "135 W 45th St, New York, NY 10036, United States",
				UserName:  "cleberson",
				Password:  "123",
			},
			userFound: model.User{
				ID:        uuid.New(),
				Name:      "Jonh Doe",
				BirthDate: time.Date(1988, 07, 07, 0, 0, 0, 0, time.UTC),
				Email:     "test@test.com",
				Address:   "135 W 45th St, New York, NY 10036, United States",
				UserName:  "cleberson",
				Password:  "123",
			},
		},
	}

	for caseTitle, tc := range cases {
		t.Run(caseTitle, func(t *testing.T) {
			repository := mocks.User{}

			repository.On(
				"Update",
				mock.Anything,
			).Return(tc.errorExpected)

			repository.On(
				"GetByID",
				mock.Anything,
			).Return(&tc.userFound, tc.errorExpected)

			userService := service.NewUserService(&repository)
			err := userService.UpdateUser(tc.user)
			assert.NoError(t, err)
		})
	}
}

func TestUpdateError(t *testing.T) {
	userID := uuid.New()
	cases := map[string]struct {
		user          domain.User
		userExpected  model.User
		errorExpected error
	}{
		"error": {
			user: domain.User{
				ID:        &userID,
				Name:      "Jonh Doe",
				BirthDate: time.Date(1988, 07, 07, 0, 0, 0, 0, time.UTC),
				Email:     "test@test.com",
				Address:   "135 W 45th St, New York, NY 10036, United States",
				UserName:  "cleberson",
				Password:  "123",
			},
			errorExpected: errors.New("user not found"),
		},
	}

	for caseTitle, tc := range cases {
		t.Run(caseTitle, func(t *testing.T) {
			repository := mocks.User{}

			repository.On(
				"Update",
				mock.Anything,
			).Return(tc.errorExpected)

			repository.On(
				"GetByID",
				mock.Anything,
			).Return(nil, tc.errorExpected)

			userService := service.NewUserService(&repository)

			err := userService.UpdateUser(tc.user)
			assert.Error(t, err)
		})
	}
}
