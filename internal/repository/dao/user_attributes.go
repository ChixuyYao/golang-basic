package dao

type UserAttribute struct {
	ID             string `gorm:"type:varchar(255);primaryKey;comment:用户属性ID(PK约束);"`
	UserID         string `gorm:"type:varchar(255);comment:用户ID(FK约束);index;not null;"`
	AttributeKey   string `gorm:"type:varchar(255);comment:属性名;index:idx_user_attr_key;"`
	AttributeValue string `gorm:"type:text;comment:属性值;"`
	ValidFrom      int64  `gorm:"type:timestamp;comment:属性生效时间;index:idx_user_attr_validity;"`
	ValidTo        int64  `gorm:"type:timestamp;comment:属性失效时间;index:idx_user_attr_validity;"`
	CreatedAt      int64  `gorm:"type:timestamp;comment:创建时间;"`
	UpdatedAt      int64  `gorm:"type:timestamp;comment:更新时间;"`
	Description    string `gorm:"type:text;comment:描述信息;"`
	//SourceSystem   string `gorm:"type:varchar(100);comment:属性来源系统"`

	// 关联关系
	//User User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"user,omitempty"`
}
