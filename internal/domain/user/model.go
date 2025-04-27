package user

import (
	"fmt"

	"github.com/mohar9h/go-graphql-ws-api/internal/domain/base"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	base.BaseModel
	Username  string `json:"username" gorm:"uniqueIndex;null"`
	Password  string `json:"password" gorm:"null"`
	Email     string `json:"email" gorm:"uniqueIndex;null"`
	Mobile    string `json:"mobile" gorm:"uniqueIndex;null"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (User) TableName() string {
	return "users"
}

func (user *User) SetPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return nil
}

func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}

func (user *User) FullName() string {
	if user.FirstName == "" && user.LastName == "" {
		return user.Username
	}
	return fmt.Sprintf("%s %s", user.FirstName, user.LastName)
}
