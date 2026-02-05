package dao

type Article struct {
	Id        string `gorm:"primaryKey"`
	Title     string `gorm:"type:varchar(4096);"`
	Content   string `gorm:"type:blob;"`
	AuthorId  string `gorm:"index"`
	CreatedAt int64  `gorm:"comment:创建时间;"`
	UpdatedAt int64  `gorm:"comment:更新时间;"`
}
