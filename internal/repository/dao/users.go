package dao

import (
	"context"
	"errors"
	"time"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type UserDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) *UserDao {
	return &UserDao{db: db}
}

var (
	ErrDuplicateEmail = errors.New("邮箱已经存在")
	ErrRecordNotFound = gorm.ErrRecordNotFound
)

func (dao *UserDao) Insert(ctx context.Context, u User) error {
	now := time.Now().UnixMilli() // 当前时间的毫秒数
	u.CTime = now
	u.UTime = now
	err := dao.db.WithContext(ctx).Create(&u).Error
	if me, ok := err.(*mysql.MySQLError); ok {
		const duplicateErr uint16 = 1062
		if me.Number == duplicateErr {
			// 邮箱冲突(用户冲突)
			return ErrDuplicateEmail
		}
	}
	return err
}

// User 数据库表的字段
type User struct {
	Id       int64  `gorm:"primary_key,auto_increment"`
	Email    string `gorm:"unique"`
	Password string

	// 创建时间,更新时间(UTC +0)
	CTime int64
	UTime int64

	// 如果需要存储JSON,则使用String类型
}
