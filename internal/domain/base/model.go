package base

import "time"

type BaseModel struct {
	ID        int64     `json:"id" gorm:"primary_key;auto_increment"`
	IsDeleted bool      `json:"is_deleted" gorm:"default:false"`
	IsActive  bool      `json:"is_active" gorm:"default:true"`
	CreatedBy string    `json:"created_by" gorm:"null"`
	UpdatedBy string    `json:"updated_by" gorm:"null"`
	DeletedBy string    `json:"deleted_by" gorm:"null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt time.Time `json:"deleted_at" gorm:"null"`
}
