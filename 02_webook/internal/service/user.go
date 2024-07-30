package service

import (
	"context"
	"geekbang/basic-go/02_webook/internal/domain"
	"geekbang/basic-go/02_webook/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

var (
	// ErrUserDuplicateEmail alias
	ErrUserDuplicateEmail = repository.ErrUserDuplicateEmail
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{
		repo: userRepository,
	}
}

// SignUp align with handler, pass ctx from for
func (svc *UserService) SignUp(ctx context.Context, user domain.User) error {
	// encryption & store
	encrypted, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(encrypted)

	// call repo
	return svc.repo.Create(ctx, user)

}
