package dao

import (
	"context"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LanguagesDao interface {
	SelectAll(ctx context.Context) ([]Languages, error)
	Insert(ctx context.Context, data Languages) error
	Update(ctx context.Context, id string, data Languages) error
	Delete(ctx context.Context, id string) error
}

// NewLanguageDao 实例化DAO对象
func NewLanguageDao(db *gorm.DB) LanguagesDao {
	return &GORMLanguagesDao{db: db}
}

// GORMLanguagesDao 基于GORM驱动的实现
type GORMLanguagesDao struct {
	db *gorm.DB
}

func (dao *GORMLanguagesDao) SelectAll(ctx context.Context) ([]Languages, error) {
	var languages []Languages
	err := dao.db.WithContext(ctx).Find(&languages).Error
	return languages, err
}

func (dao *GORMLanguagesDao) Insert(ctx context.Context, data Languages) error {
	now := time.Now().UnixMilli()
	data.Id = uuid.New().String()
	data.CreatedAt = now
	data.UpdatedAt = now
	err := dao.db.WithContext(ctx).Create(&data).Error
	return err
}

func (dao *GORMLanguagesDao) Update(ctx context.Context, id string, data Languages) error {
	now := time.Now().UnixMilli()
	data.UpdatedAt = now
	err := dao.db.WithContext(ctx).Where("id = ?", id).Updates(data).Error
	return err
}

func (dao *GORMLanguagesDao) Delete(ctx context.Context, id string) error {
	return dao.db.WithContext(ctx).Delete(&Languages{}, id).Error
}

// Languages 语言表
type Languages struct {
	Id        string `gorm:"primaryKey;comment:语言唯一标识符，自增主键"`
	ISO       string `gorm:"type:varchar(10);uniqueIndex;comment:ISO 639-1/639-2标准语言代码，必须唯一"`
	AliasCn   string `gorm:"type:varchar(100);comment:英文名称"`
	AliasEn   string `gorm:"type:varchar(100);comment:中文名称"`
	Name      string `gorm:"type:varchar(100);comment:本族名称，支持原生字符显示"`
	Script    string `gorm:"type:varchar(50);comment:文字系统，如拉丁字母、汉字等"`
	Direction string `gorm:"comment:书写方向, enum(ltr=左到右,rtl=右到左)"`
	Family    string `gorm:"comment:语言所属语系，用于分类"`
	IsActive  bool   `gorm:"comment:是否激活使用，TRUE=激活，FALSE=停用"`
	CreatedAt int64  `gorm:"comment:创建时间"`
	UpdatedAt int64  `gorm:"comment:更新时间"`
}
