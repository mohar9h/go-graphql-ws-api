package permission

import (
	"github.com/mohar9h/go-graphql-ws-api/internal/domain/base"
)

type Permission struct {
	base.BaseModel
	UserID   *uint  `json:"user_id" gorm:"index"`
	GroupID  *uint  `json:"group_id" gorm:"index"`
	Resource string `json:"resource" gorm:"size:50;index;not null"`
	Bitmask  uint64 `json:"bitmask" gorm:"not null"`
}

func (Permission) TableName() string {
	return "permissions"
}
