package user

import (
	"commerce/internal/shared/models"
	"time"

	"gorm.io/gorm"
)

type UserRepositoryI interface {
	GetById(id uint) (*models.User, error)
	GetAll() ([]*models.User, error)
	Save(user *models.User) error
	Delete(id uint, hard bool) error
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepositoryI {
	return &UserRepository{db: db}
}

// Delete implements [UserRepositoryI].
func (u *UserRepository) Delete(id uint, hard bool) error {
	if hard {
		return u.db.Delete(models.User{}, id).Error
	}
	var user *models.User
	if err := u.db.First(&user, id).Error; err != nil {
		return err
	}
	user.DeletedDate = time.Now()
	return u.db.Save(user).Error
}

// GetAll implements [UserRepositoryI].
func (u *UserRepository) GetAll() ([]*models.User, error) {
	var users []*models.User
	if err := u.db.Find(models.User{}, &users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// GetById implements [UserRepositoryI].
func (u *UserRepository) GetById(id uint) (*models.User, error) {
	var user models.User
	if err := u.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// Save implements [UserRepositoryI].
func (u *UserRepository) Save(user *models.User) error {
	if user.Id == 0 {
		return u.db.Create(user).Error
	} else {
		if _, err := u.GetById(user.Id); err != nil {
			return err
		}
		return u.db.Save(user).Error
	}
}
