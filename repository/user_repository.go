package repository

import (
	"time"

	"vigour/model"
	"vigour/infrastructure"
) 

type userRepository struct {}

func (r *userRepository) GetAll() ([]model.User, error) {
	db := infrastructure.GetDB()
	var users []model.User

	if err := db.Model(&model.User{}).Preload("Profile").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) GetById(id int) (*model.User, error) {
	db := infrastructure.GetDB()
	var user model.User
	if err := db.Preload("Profile").First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) CreateUser(newUser *model.User) (*model.User, error) {
	db := infrastructure.GetDB()
	if err := db.Create(newUser).Error; err != nil {
		return nil, err
	}
	return newUser, nil
}

func (r *userRepository) GetByUsername(username string) (*model.User, error) {
	db := infrastructure.GetDB()

	var user model.User
	if err := db.Model(&model.User{}).Where("username = ?", username).Find(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) UpdateUser(newUser model.User) (*model.User, error) {
	db := infrastructure.GetDB()
	
	if err := db.Model(&newUser).Where("id = ?", newUser.Id).Update("password",newUser.Password).Error; err != nil {
		return nil, err
	}
	return &newUser, nil
}

func (r *userRepository) DeleteUser(id int) (*model.User, error) {
	db := infrastructure.GetDB()

	if err := db.Model(&model.User{Id: id}).Update("deleted_at", time.Now()).Error; err != nil {
		return nil, err
	}

	return r.GetById(id)
}

func (r *userRepository) LoginTokenRequest(user *model.User) (bool, error) {
	db := infrastructure.GetDB()

	var userInfo model.User
	if err := db.Where(&model.User{
		Username: user.Username,
		Password: user.Password,
	}).Find(&userInfo).Error; err != nil {
		return false, nil
	}

	user.ExpiresAt = time.Now().Local().Add(time.Hour*time.Duration(infrastructure.Extend_Hour)).UnixNano() / infrastructure.NANO_TO_SECOND
	return true, nil
}


func NewUserRepository() model.UserRepository {
	return &userRepository{}
}
