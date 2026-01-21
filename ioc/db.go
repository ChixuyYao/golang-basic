package ioc

import (
	"golang/config"
	rbacDao "golang/internal/repository/dao"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open(config.BookConfig.DB.DSN), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, // 禁用数据库外键约束生成
	})
	if err != nil {
		panic(err) //数据库启动失败,终止进程
	}
	err = rbacDao.InitTables(db)
	if err != nil {
		panic(err) //数据库表创建失败,终止进程
	}
	return db
}
