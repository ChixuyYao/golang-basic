package dao

import (
	"context"
	"database/sql"
	"errors"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

func NewUserDao(db *gorm.DB) UserDao {
	return &GORMUsersDao{db: db}
}

// UserDao 登录(查询),注册(插入),更改密码,邮箱,手机,账户状态(修改),注销账户(删除)
type UserDao interface {
	Select(ctx context.Context, entity Users) (Users, error)
	Insert(ctx context.Context, entity Users) error
	update(ctx context.Context, id string, entity Users) error
	Delete(ctx context.Context, id string) error
}

var (
	ErrDuplicate      = errors.New("邮箱,手机号码存在冲突")
	ErrRecordNotFound = gorm.ErrRecordNotFound
)

type GORMUsersDao struct {
	db *gorm.DB
}

// Select 查询USER数据(EX:用户登录)
func (dao *GORMUsersDao) Select(ctx context.Context, entity Users) (Users, error) {
	var user Users
	err := dao.db.WithContext(ctx).
		//Where(&Users{Password: entity.Password}).
		Or(&Users{Email: entity.Email, Phone: entity.Phone}).
		First(&user).
		Error
	return user, err
}

// Insert 插入USER数据(EX:用户注册)
func (dao *GORMUsersDao) Insert(ctx context.Context, entity Users) error {

	err := dao.db.WithContext(ctx).Create(&entity).Error
	if me, ok := err.(*mysql.MySQLError); ok {
		const duplicateErr uint16 = 1062 // 数据冲突
		if me.Number == duplicateErr {
			return ErrDuplicate
		}

	}
	return err
}

func (dao *GORMUsersDao) update(ctx context.Context, id string, entity Users) error {
	//TODO implement me
	return nil
}

func (dao *GORMUsersDao) Delete(ctx context.Context, id string) error {
	//TODO implement me
	return nil
}

// Users ABAC数据表预留
type Users struct {
	Id        string         `gorm:"type:varchar(255);primaryKey;comment:用户ID(PK);"`
	Username  string         `gorm:"type:varchar(255);comment:用户昵称;"`
	Password  string         `gorm:"comment:用户密码(加密);not null;"`
	Email     sql.NullString `gorm:"type:varchar(255);comment:邮箱账户(用于登录);uniqueIndex;"`
	Phone     sql.NullString `gorm:"type:varchar(128);comment:手机号码(用于登录);uniqueIndex;"`
	State     int            `gorm:"comment:账户状态(0=正常,1=冻结,2=停用);"`
	CreatedAt int64          `gorm:"comment:创建时间;"`
	UpdatedAt int64          `gorm:"comment:更新时间;"`
	LastLogin int64          `gorm:"comment:最后登录时间;"`
	// 关联关系
	//Attributes []UserAttribute         `gorm:"foreignKey:UserID" json:"attributes,omitempty"`
	//AccessLogs []AccessLog            `gorm:"foreignKey:UserID" json:"access_logs,omitempty"`
	//Contexts   []EnvironmentContext   `gorm:"foreignKey:UserID" json:"contexts,omitempty"`
	//Policies   []Policy               `gorm:"many2many:policy_targets;" json:"policies,omitempty"`
}
