package group

import (
	"github.com/mohar9h/go-graphql-ws-api/internal/domain/base"
)

type Group struct {
	base.BaseModel
	Name        string `json:"name" gorm:"size:50;uniqueIndex;not null"`
	Description string `json:"description" gorm:"size:255"`
}

func (Group) TableName() string {
	return "groups"
}
