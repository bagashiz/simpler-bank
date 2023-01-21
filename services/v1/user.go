package v1service

import (
	"context"

	"github.com/bagashiz/simpler-bank/db"
	"github.com/bagashiz/simpler-bank/models"
)

// UserService is a struct for working with user related database operations.
type UserService struct {
	User models.User
}

// CreateUser is a function for creating a new user.
func (us *UserService) CreateUser(ctx context.Context) (models.User, error) {
	if err := db.GetDB().WithContext(ctx).Create(&us.User).Error; err != nil {
		return us.User, err
	}

	return us.User, nil
}
