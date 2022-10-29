package model

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type User struct {
	Id                 int        `json:"id" gorm:"primaryKey"`
	Username           string     `json:"username" gorm:"column:username;unique"`
	Password           string     `json:"password" gorm:"password"`
	Role               string     `json:"role" gorm:"role"`
	Profile            *Profile   `json:"profile" gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	DeletedAt          *time.Time `json:"deleted_at" gorm:"deleted_at"`
	jwt.StandardClaims `gorm:"-" swaggerignore:"true"`
}

type UserPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserResponse struct {
	Id       int      `json:"id" gorm:"primaryKey"`
	Username string   `json:"username" gorm:"unique;column:username"`
	Role     string   `json:"role" gorm:"column:role"`
	Profile  *Profile `json:"profile"`
}

type UserRepository interface {
	GetAll() ([]User, error)
	CreateUser(newUser *User) (*User, error)
	GetById(id int) (*User, error)
	GetByUsername(username string) (*User, error)
	UpdateUser(newUser User) (*User, error)
	DeleteUser(id int) (*User, error)
	LoginTokenRequest(*User) (bool, error)
}
