package dao

import (
	"gorm.io/gorm"
)

type LanguagesDao struct {
	db *gorm.DB
}

func NewLanguageDao(db *gorm.DB) *LanguagesDao {
	return &LanguagesDao{db: db}
}

// Languages 语言表(V2)
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
