package services

import (
	"errors"
	"stockhive-server/internal/config"
	"stockhive-server/internal/models"
	"stockhive-server/internal/repositories"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService interface {
	Register(user *models.User) error
	Login(username, password string) (string, string, error)
	RefreshToken(token string) (string, error)
	ChangePassword(req *models.ChangePasswordRequest) error
	GetAllUsers() ([]models.User, error)
	GetUserByID(id string) (models.User, error)
	UpdateUser(id string, updated models.User) (models.User, error)
	DeleteUser(id string) error

}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) UpdateUser(id string, updated models.User) (models.User, error){
	user, err := s.repo.FindByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return models.User{}, errors.New("user not found")
		}
		return models.User{}, err
	}

	err = s.repo.Update(&user, updated)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (s *userService) DeleteUser(id string) error {
	user, err := s.repo.FindByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("user not found")
		}
	}

	return s.repo.Delete(&user)
}


func (s *userService) GetUserByID(id string) (models.User, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return models.User{}, errors.New("user not found")
		}
		return models.User{}, err
	}
	return user, nil
}

func (s *userService) GetAllUsers() ([]models.User, error) {
	users, err := s.repo.GetAllUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *userService) Register(user *models.User) error {
	// Cek apakah user sudah ada
	// _, err := s.repo.FindByUsernameOrEmail(user.Username, user.Email)
	// if err == nil  {
	// 	return errors.New("username or email already registered")
	// }
	// if err == nil && user.Email == "" {
	// 	return errors.New("username already registered")
	// }

	// Cek apakah email sudah ada
	if user.Email != "" {
		_, err := s.repo.FindByEmail(user.Email)
		if err == nil {
			return errors.New("email already registered")
		}
	}

	//cek apakah username sudah ada
	_, err := s.repo.FindByUsername(user.Username)
	if err == nil {
		return errors.New("username already registered")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	// Simpan user
	return s.repo.CreateUser(user)
}

func (s *userService) Login(username, password string) (string, string, error) {
	user, err := s.repo.FindByUsername(username)
	if err != nil {
		return "", "", errors.New("username not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", "", errors.New("invalid password")
	}

	accessToken, err := config.GenerateToken(user.Username, 15*time.Minute, "access")
	if err != nil {
		return "", "", err
	}

	refreshToken, err := config.GenerateToken(user.Username, 168*time.Hour, "refresh")
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (s *userService) RefreshToken(token string) (string, error) {
	claims, err := config.VerifyToken(token, "refresh")
	if err != nil {
		return "", errors.New("invalid refresh token")
	}

	username := (*claims)["username"].(string)
	newAccessToken, err := config.GenerateToken(username, 15*time.Minute, "access")
	if err != nil {
		return "", err
	}

	return newAccessToken, nil
}

func (s *userService) ChangePassword(req *models.ChangePasswordRequest) error {
	user, err := s.repo.FindByUsername(req.Username)
	if err != nil {
		return errors.New("username not found")
	}

	// Verifikasi password lama
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.CurrentPassword)); err != nil {
		return errors.New("invalid current password")
	}

	// Pastikan password baru sesuai dengan konfirmasi password
	if req.NewPassword != req.ConfirmPassword {
		return errors.New("new password and confirm password do not match")
	}

	// Hash password baru
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Update password di database
	return s.repo.UpdatePassword(user, string(hashedPassword))
}