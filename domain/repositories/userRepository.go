package repositories

import (
	"auth/database"
	"auth/domain/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindByEmail(email string) (*models.User, error)
	CreateUser(m *models.User) (*models.User, error)
}

type UserRepo struct {
	DB *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *UserRepo {
	return &UserRepo{
		DB: db,
	}
}

func (r *UserRepo) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.DB.First(&user, "email = ?", email).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (r *UserRepo) CreateUser(user *models.User) (*models.User, error) {
	result := database.DB.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

//func (r *UserRepo) CreateUser(email string, password string) (*models.User, error) {
//	var user models.User
//	err := r.DB.Create(&user, "email = ?", email).Error
//	if err != nil {
//		return nil, err
//	}
//	return &user, nil
//}
