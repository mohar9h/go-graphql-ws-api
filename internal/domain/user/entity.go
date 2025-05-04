package user

import (
	"fmt"
	"github.com/mohar9h/go-graphql-ws-api/internal/domain/group"

	"github.com/mohar9h/go-graphql-ws-api/internal/domain/base"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	base.BaseModel
	Username  string      `json:"username" gorm:"uniqueIndex;size:20"`
	Password  string      `json:"password" gorm:"size:255"`
	Email     string      `json:"email" gorm:"uniqueIndex;size:100"`
	Mobile    string      `json:"mobile" gorm:"uniqueIndex;size:11"`
	FirstName string      `json:"first_name" gorm:"size:50"`
	LastName  string      `json:"last_name" gorm:"size:50"`
	GroupId   int64       `json:"group_id" gorm:"index"`
	Group     group.Group `json:"group" gorm:"foreignkey:GroupId"`
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
