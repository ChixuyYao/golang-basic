package dao

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	mysql2 "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestGORMUsersDao_Insert(t *testing.T) {
	// 测试用例切片
	testcases := []struct {
		name string // 测试用例名称
		mock func(t *testing.T) *sql.DB
		ctx  context.Context
		user Users

		wantErr error
	}{
		{
			name: "数据插入成功",
			mock: func(t *testing.T) *sql.DB {
				db, mock, err := sqlmock.New()
				assert.NoError(t, err)
				mockResult := sqlmock.NewResult(1, 1)
				mock.ExpectExec("INSERT INTO .*").
					WillReturnResult(mockResult)
				return db
			},
			ctx:  context.Background(),
			user: Users{Username: "Testing Username"},
		},
		{
			name: "数据插入失败:邮箱,手机号冲突",
			mock: func(t *testing.T) *sql.DB {
				db, mock, err := sqlmock.New()
				assert.NoError(t, err)
				mock.ExpectExec("INSERT INTO .*").
					WillReturnError(&mysql2.MySQLError{Number: 1062})
				return db
			},
			ctx:     context.Background(),
			user:    Users{Username: "Testing Username"},
			wantErr: ErrDuplicate,
		},
		{
			name: "数据插入失败:数据库异常,错误",
			mock: func(t *testing.T) *sql.DB {
				db, mock, err := sqlmock.New()
				assert.NoError(t, err)
				mock.ExpectExec("INSERT INTO .*").
					WillReturnError(errors.New("数据库异常"))
				return db
			},
			ctx:     context.Background(),
			user:    Users{Username: "Testing Username"},
			wantErr: errors.New("数据库异常"),
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			sqlDB := tc.mock(t)
			db, err := gorm.Open(mysql.New(mysql.Config{
				Conn:                      sqlDB,
				SkipInitializeWithVersion: true,
			}), &gorm.Config{
				DisableAutomaticPing:   true, // 关闭数据库自动重连
				SkipDefaultTransaction: true, // 关闭数据库事务自动提交
			})
			assert.NoError(t, err)
			dao := NewUserDao(db)
			err = dao.Insert(tc.ctx, tc.user)
			assert.Equal(t, tc.wantErr, err)
		})
	}
}
