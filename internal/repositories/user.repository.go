package repositories

import (
	"stockhive-server/internal/config"
	"stockhive-server/internal/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	// FindByUsernameOrEmail(username, email string) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	CreateUser(user *models.User) error
	FindByUsername(username string) (*models.User, error)
	UpdatePassword(user *models.User, newPassword string) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository() UserRepository {
	return &userRepository{
		db: config.ConnectDB(),
	}
}

// func (r *userRepository) FindByUsernameOrEmail(username, email string) (*models.User, error) {
// 	var user models.User
// 	err := r.db.Where("username = ? OR email = ?", username, email).First(&user).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &user, nil
// }

// func (r *userRepository) FindByUsername(username string) (*models.User, error) {
// 	var user models.User
// 	err := r.db.Where("username = ?", username).First(&user).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &user, nil
// }

func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User

	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) UpdatePassword(user *models.User, newPassword string) error {
	return r.db.Model(user).Update("password", newPassword).Error
}