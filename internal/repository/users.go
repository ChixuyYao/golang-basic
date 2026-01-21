package repository

import (
	"context"
	"database/sql"
	"golang/internal/domain"
	"golang/internal/repository/dao"
	"time"
)

func NewUserRepository(dao dao.UserDao) UserRepository {
	return &userRepository{
		dao: dao,
	}
}

type UserRepository interface {
	FindByEmailOrPhone(ctx context.Context, email string, phone string) (domain.Users, error)
	Create(ctx context.Context, user domain.Users) error
}

var (
	ErrDuplicate    = dao.ErrDuplicate
	ErrUserNotFound = dao.ErrRecordNotFound
)

type userRepository struct {
	dao dao.UserDao
}

// FindByEmailOrPhone 根据邮箱,手机号码查询用户
func (repo *userRepository) FindByEmailOrPhone(ctx context.Context, email, phone string) (domain.Users, error) {
	user, err := repo.dao.Select(ctx, repo.toEntity(domain.Users{Email: email, Phone: phone}))
	if err != nil {
		return domain.Users{}, err
	}
	return repo.toDomain(user), nil
}

// Create 创建新的用户
func (repo *userRepository) Create(ctx context.Context, user domain.Users) error {
	return repo.dao.Insert(ctx, repo.toEntity(user))
}

// 数据库对象 -> 领域对象
func (repo *userRepository) toDomain(u dao.Users) domain.Users {
	return domain.Users{
		Id:        u.Id,
		Email:     u.Email.String,
		Phone:     u.Phone.String,
		Username:  u.Username,
		Password:  u.Password,
		CreatedAt: time.UnixMilli(u.CreatedAt),
		UpdatedAt: time.UnixMilli(u.UpdatedAt),
	}
}

// 领域对象 -> 数据库对象
func (repo *userRepository) toEntity(u domain.Users) dao.Users {
	return dao.Users{
		Id:       u.Id,
		Username: u.Username,
		Password: u.Password,
		Email:    sql.NullString{String: u.Email, Valid: u.Email != ""},
		Phone:    sql.NullString{String: u.Phone, Valid: u.Phone != ""},
	}
}
