package repositories

import (
	"go-fiber/data/output"
	"go-fiber/domain/entities"

	"gorm.io/gorm"
)


type UserRepository struct {
	db *gorm.DB
}

// GetAllUsers implements UserRepository.
func (u *UserRepository) GetAllUsers() {
	panic("unimplemented")
}
func (u *UserRepository) CreateUser(data entities.UserEntity) error {
	tx := u.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Create(&data).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}


func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

var _ output.UserRepository = (*UserRepository)(nil)
