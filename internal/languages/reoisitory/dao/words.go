package dao

import (
	"gorm.io/gorm"
)

type WordsDao struct {
	db *gorm.DB
}

func NewWordsDao(db *gorm.DB) *WordsDao {
	return &WordsDao{db: db}
}

// Words 词汇表(主表)
type Words struct {
	Id            string `gorm:"primaryKey;comment:词汇唯一标识符，自增主键"`
	LanguageID    string `gorm:"index:idx_language;comment:语言ID"`
	PosID         string `gorm:"index:idx_pos;comment:词性ID"`
	Word          string `gorm:"comment:词汇"`
	Base          string `gorm:"comment:词汇原形(标准形式)"`
	Ipa           string `gorm:"comment:国际音标(IPA)"`
	TranslationCn string `gorm:"comment:中文释义"`
	TranslationEn string `gorm:"comment:英文释义"`
	Etymology     string `gorm:"comment:词源信息"`
	CEFR          string `gorm:"comment:欧标等级 CEFR:A1-C2"`
	IsCommon      bool   `gorm:"comment:是否常用词，TRUE=常用，FALSE=不常用"`
	FrequencyRank int    `gorm:"comment:使用频率，数值越小越常用，NULL=未统计"`
	CreatedAt     int64  `gorm:"comment:创建时间"`
	UpdatedAt     int64  `gorm:"comment:更新时间"`

	// 关联关系
	//Language        Languages      `gorm:"foreignKey:LanguageID;references:LanguageID;constraint:OnDelete:CASCADE;comment:所属语言"`
	//PartOfSpeech    PartOfSpeech  `gorm:"foreignKey:PosID;references:PosID;constraint:OnDelete:RESTRICT;comment:词性"`
}
