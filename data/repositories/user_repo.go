package repositories

import "gorm.io/gorm"

type UserRepository interface {
	//Methods
	GetAllUsers()
}

type userRepository struct {
	db *gorm.DB
}

// GetAllUsers implements UserRepository.
func (u *userRepository) GetAllUsers() {
	panic("unimplemented")
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}
