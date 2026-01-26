package dao

import (
	"context"
	"database/sql"

	"gorm.io/gorm"
)

// NewCategoriesDao 事项分类表
func NewCategoriesDao(db *gorm.DB) CategoriesDao {
	return &GORMCategoryDao{db: db}
}

type CategoriesDao interface {
	SelectByUid(ctx context.Context, uid string) ([]Categories, error)
	InsertWithUid(ctx context.Context, uid string, entity Categories) error
	Update(ctx context.Context, id string, entity Categories) error
	Delete(ctx context.Context, id string) error
}

type GORMCategoryDao struct {
	db *gorm.DB
}

// SelectByUid 查询全部分类列表(根据UserId)
func (dao *GORMCategoryDao) SelectByUid(ctx context.Context, uid string) ([]Categories, error) {
	var category []Categories
	err := dao.db.WithContext(ctx).
		Where(map[string]interface{}{"user_id": uid}).
		Find(&category).
		Error

	return category, err
}

// InsertWithUid 插入分类列表
func (dao *GORMCategoryDao) InsertWithUid(ctx context.Context, uid string, entity Categories) error {
	entity.UserId = uid
	err := dao.db.WithContext(ctx).Create(&entity).Error
	return err
}

// Update 更新分类列表
func (dao *GORMCategoryDao) Update(ctx context.Context, id string, entity Categories) error {
	entity.Id = id
	err := dao.db.WithContext(ctx).Updates(&entity).Error
	//err := dao.db.WithContext(ctx).Model(&entity).
	//	Where(map[string]interface{}{"id": id}).
	//	Updates(map[string]interface{}{}).Error
	return err
}

// Delete 删除单项分类
func (dao *GORMCategoryDao) Delete(ctx context.Context, id string) error {
	err := dao.db.WithContext(ctx).Delete(map[string]interface{}{"id": id}).Error
	return err
}

// Categories 事项分类表[结构]
type Categories struct {
	Id          string         `gorm:"type:varchar(255);comment:分类ID(PK);primaryKey;"`
	UserId      string         `gorm:"type:varchar(255);comment:用户ID(FK);"`
	Name        sql.NullString `gorm:"type:varchar(255);comment:分类名称;"`
	Description string         `gorm:"comment:分类备注;"`
	Color       string         `gorm:"comment:分类颜色(HEX);"`
	Icon        string         `gorm:"comment:分类图标;"`
	Type        int            `gorm:"分类类型(0=系统内置,1=用户创建)"`
	SortOrder   int            `gorm:"comment:排序信息;"`
	IsActive    int            `gorm:"comment:活跃状态(0=非活跃,1=活跃);"`
	CreatedAt   int64          `gorm:"comment:创建时间;"`
	UpdatedAt   int64          `gorm:"comment:更新时间;"`
}

/*
字段名			数据类型			约束						默认值				示例数据
id				BIGINT			PK, AUTO_INCREMENT							1
user_id			BIGINT			FK(users.id)								1
name			VARCHAR(100)	NOT NULL									"工作"
description		TEXT			NULL										"工作任务管理"
color			VARCHAR(7)													"#3B82F6"
icon			VARCHAR(50)		NULL										"briefcase"
sort_order		INT										0					1
is_active		BOOLEAN									TRUE				TRUE
created_at		TIMESTAMP								CURRENT_TIMESTAMP	"2024-01-15 10:00:00"
updated_at		TIMESTAMP								CURRENT_TIMESTAMP	"2024-01-15 10:00:00"

*/
