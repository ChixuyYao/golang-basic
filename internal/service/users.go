package service

import (
	"context"
	"errors"
	"golang/internal/domain"
	"golang/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

type UserService interface {
	UserLogin(ctx context.Context, email string, phone string, password string) (domain.Users, error)
	UserSignup(ctx context.Context, user domain.Users) error
}

var (
	ErrDuplicateEmail        = repository.ErrDuplicate
	ErrInvalidUserOrPassword = errors.New("账户名称,密码无效")
)

type userService struct {
	repo repository.UserRepository
}

// UserLogin 用户登录
func (svc *userService) UserLogin(ctx context.Context, email, phone, password string) (domain.Users, error) {
	user, err := svc.repo.FindByEmailOrPhone(ctx, email, phone)
	switch err {
	case nil:
	case repository.ErrUserNotFound:
		return domain.Users{}, ErrInvalidUserOrPassword
	default:
		return domain.Users{}, err
	}
	// 密码校验
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return domain.Users{}, ErrInvalidUserOrPassword
	}
	return user, nil
}

// UserSignup 注册用户
func (svc *userService) UserSignup(ctx context.Context, user domain.Users) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil { // 密码加密异常
		return err
	}
	user.Password = string(hash)
	return svc.repo.Create(ctx, user)
}
