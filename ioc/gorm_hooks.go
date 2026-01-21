package ioc

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GormHooks struct {
}

func NewGormHooks() *GormHooks {
	return &GormHooks{}
}

func (hook *GormHooks) InitHooks(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		db.Callback().Create().Before("gorm:create").Register("global:before_create", hook.beforeCreate)
	}
}

func (hook *GormHooks) beforeCreate(db *gorm.DB) {
	// 准备插入数据,生成ID(UUID),变更CreatedAt(int64),UpdateAt(int64)
	if field := db.Statement.Schema.LookUpField("Id"); field != nil {
		db.Statement.SetColumn("id", uuid.New().String())
	}
	timestamp := time.Now()
	if field := db.Statement.Schema.LookUpField("CreatedAt"); field != nil {
		db.Statement.SetColumn("created_at", timestamp)
	}
	if field := db.Statement.Schema.LookUpField("updated_at"); field != nil {
		db.Statement.SetColumn("updated_at", timestamp)
	}
}
