package users

import (
	utils2 "bridge/users-service/pkg/utils"
	"errors"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func ProvideUserRepository(DB *gorm.DB) *UserRepository {
	return &UserRepository{DB: DB}
}

func (u *UserRepository) Find(id uint) (User, error) {
	var (
		user User
		err  error
	)
	result := u.DB.First(&user, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		err = ErrUserNotFound
	}
	if result.Error != nil {
		err = utils2.WrapErrors(ErrUserUnhandledError, err)
	}
	return user, err
}

func (u *UserRepository) Get(filter *User) ([]User, error) {
	var (
		users []User
		err   error
	)
	result := u.DB.Where(filter).Find(&users)
	if result.Error != nil {
		err = utils2.WrapErrors(ErrUserUnhandledError, err)
	}
	return users, err
}

func (u *UserRepository) Create(dto CreateUserDTO) (User, error) {
	var (
		user User
		err  error
	)
	existingUsers, err := u.Get(&User{Email: dto.Email})
	if err != nil {
		return user, err
	}
	if len(existingUsers) > 0 {
		return user, ErrUserAlreadyExists
	}
	encryptedPassword, err := utils2.HashPassword(dto.Password)
	if err != nil {
		return User{}, utils2.WrapErrors(ErrUserUnhandledError, err)
	}
	user = User{Email: dto.Email, Password: encryptedPassword}
	result := u.DB.Create(&user)

	if result.Error != nil {
		return user, utils2.WrapErrors(ErrUserUnhandledError, err)
	}

	return user, result.Error
}

func (u *UserRepository) Update(id uint, dto UpdateUserDTO) (User, error) {
	var (
		user User
		err  error
	)

	user, err = u.Find(id)
	if err != nil {
		return user, err
	}

	if dto.Email != "" {
		user.Email = dto.Email
	}

	if dto.Password != "" {
		encryptedPassword, err := utils2.HashPassword(dto.Password)
		if err != nil {
			return User{}, utils2.WrapErrors(ErrUserUnhandledError, err)
		}
		user.Password = encryptedPassword
	}

	result := u.DB.Save(&user)

	if result.Error != nil {
		return user, utils2.WrapErrors(ErrUserUnhandledError, result.Error)
	}

	return user, err
}

func (u *UserRepository) Delete(id uint) (int64, error) {
	var err error
	result := u.DB.Delete(&User{}, id)
	if result.Error != nil {
		err = utils2.WrapErrors(ErrUserUnhandledError, err)
	}
	return result.RowsAffected, err
}
