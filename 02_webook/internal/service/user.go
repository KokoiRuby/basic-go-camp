package service

import (
	"context"
	"errors"
	"geekbang/basic-go/02_webook/internal/domain"
	"geekbang/basic-go/02_webook/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

var (
	// ErrUserDuplicateEmail alias
	ErrUserDuplicateEmail     = repository.ErrUserDuplicateEmail
	ErrInvalidEmailOrPassword = errors.New("invalid email or password")
	ErrUserNotFound           = repository.ErrUserNotFound
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

func (svc *UserService) Login(ctx context.Context, user domain.User) (domain.User, error) {
	// 1. user find by email
	u, err := svc.repo.FindByEmail(ctx, user)
	if errors.Is(err, ErrUserNotFound) {
		return domain.User{}, ErrInvalidEmailOrPassword
	}
	if err != nil {
		return domain.User{}, err
	}
	// 2. compare pwd
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(user.Password))
	if err != nil {
		return domain.User{}, ErrInvalidEmailOrPassword
	}
	return u, nil
}
