package dao

import (
	"gorm.io/gorm"
)

type PartOfSpeechDao struct {
	db *gorm.DB
}

func NewPartOfSpeechDao(db *gorm.DB) *PartOfSpeechDao {
	return &PartOfSpeechDao{}
}

// PartOfSpeech 词性信息表
type PartOfSpeech struct {
	Id           string `gorm:"primaryKey;comment:词性唯一标识符，自增主键"`
	LanguageID   string `gorm:"comment:语言ID"`
	Name         string `gorm:"comment:词性全量名称，目标语言术语"`
	Abbreviation string `gorm:"comment:词性缩写名称，用于显示查询"`
	Description  string `gorm:"type:text;comment:词性详细描述，包括用法和示例"`
	ParentId     string `gorm:"comment:父词性ID，用于构建词性层级结构，NULL表示顶级词性"`
	Sorted       int    `gorm:"comment:排序顺序，数值越小越靠前"`
	CreatedAt    int64  `gorm:"comment:创建时间"`
	UpdatedAt    int64  `gorm:"comment:更新时间"`

	// 关联关系
	//Language      Languages       `gorm:"foreignKey:LanguageID;references:LanguageID;constraint:OnDelete:CASCADE;comment:所属语言"`
	//Parent        *PartOfSpeech  `gorm:"foreignKey:ParentPosID;references:PosID;constraint:OnDelete:SET NULL;comment:父词性"`
	//Children      []PartOfSpeech `gorm:"foreignKey:ParentPosID;references:PosID;comment:子词性列表"`
}
