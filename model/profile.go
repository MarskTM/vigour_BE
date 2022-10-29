package model

import (
	"time"
)

type Profile struct {
	Id       int       `json:"id" gorm:"primaryKey"`
	RealName string    `json:"real_name"`
	Name     string    `json:"name"`
	Birthday time.Time `json:"birthday"`
	Gender   string    `json:"gender"`
	Email    string    `json:"email"`
	Phone    string    `json:"phone"`
	Avatar   string    `json:"avatar"`
	NewUser  string    `json:"newUser"`

	UserId int `json:"userId" gorm:"column:user_id;unique"`

	CreatedAt time.Time  `swaggerignore:"true"`
	UpdatedAt time.Time  `swaggerignore:"true"`
	DeletedAt *time.Time `swaggerignore:"true"`
}

type ProfileRepository interface {
	GetById(id int) (*Profile, error)
	GetAll() ([]Profile, error)
	GetByUserId(user_id int) (*Profile, error)
	Create(new *Profile) (*Profile, error)
	Update(id int, profile Profile) (*Profile, error)
	Delete(id int) error
	Upsert(profile *Profile) (*Profile, error)
}
