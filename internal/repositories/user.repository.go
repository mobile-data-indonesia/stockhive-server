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
	GetAllUsers() ([]models.User, error)
	FindByID(id string) (models.User, error)
	Delete(user *models.User) error
	Update(user *models.User, updated models.User) error
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

//repo to delete
func (r *userRepository) Delete(user *models.User) error {
	return r.db.Delete(user).Error
}

func (r *userRepository) Update(user *models.User, updated models.User) error {
	return r.db.Model(user).Updates(updated).Error
}

func (r *userRepository) FindByID(id string) (models.User, error) {
	var user models.User
	err := r.db.Where("user_id = ?", id).First(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *userRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

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