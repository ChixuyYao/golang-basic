package dao

import (
	"gorm.io/gorm"
)

func InitTables(db *gorm.DB) error {
	return db.AutoMigrate(
		//&Languages{},
		//&PartOfSpeech{},
		//&Words{},

		&Users{},
		&UserAttribute{},

		// Reminder 模块提供数据库表
		&Categories{},
		&Goals{},
		&Tasks{},
	)
}
