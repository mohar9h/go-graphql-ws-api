package user

import (
	"context"

	"gorm.io/gorm"

	"github.com/mohar9h/go-graphql-ws-api/internal/domain/user"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) user.Repository {
	return &repository{db: db}
}

func (r *repository) Count(ctx context.Context) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&user.User{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *repository) Create(ctx context.Context, user *user.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *repository) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&user.User{}, id).Error
}

func (r *repository) GetByEmail(ctx context.Context, email string) (*user.User, error) {
	var u user.User
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&u).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *repository) GetByID(ctx context.Context, id int64) (*user.User, error) {
	var u user.User
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&u).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *repository) GetByMobile(ctx context.Context, mobile string) (*user.User, error) {
	var u user.User
	err := r.db.WithContext(ctx).Where("mobile = ?", mobile).First(&u).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *repository) GetByUsername(ctx context.Context, username string) (*user.User, error) {
	var u user.User
	err := r.db.WithContext(ctx).Where("username = ?", username).First(&u).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *repository) List(ctx context.Context, offset int, limit int) ([]*user.User, int64, error) {
	var user []*user.User
	err := r.db.WithContext(ctx).Offset(offset).Limit(limit).Find(&user).Error
	if err != nil {
		return nil, 0, err
	}
	return user, int64(len(user)), nil
}

// Update implements user.Repository.
func (r *repository) Update(ctx context.Context, user *user.User) error {
	return r.db.WithContext(ctx).Save(user).Error
}
