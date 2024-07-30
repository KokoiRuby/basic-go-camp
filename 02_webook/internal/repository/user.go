package repository

import (
	"context"
	"geekbang/basic-go/02_webook/internal/domain"
	"geekbang/basic-go/02_webook/internal/repository/dao"
)

var (
	// ErrUserDuplicateEmail alias
	ErrUserDuplicateEmail = dao.ErrUserDuplicateEmail
)

type UserRepository struct {
	dao *dao.UserDAO
}

func NewUserRepository(dao *dao.UserDAO) *UserRepository {
	return &UserRepository{dao: dao}
}

func (ur *UserRepository) Create(ctx context.Context, user domain.User) error {
	return ur.dao.Insert(ctx, dao.User{
		Email:    user.Email,
		Password: user.Password,
	})
}

func (ur *UserRepository) FindById(int64) {
	// cache
	// dao if miss & write to cache

}
