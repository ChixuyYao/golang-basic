package domain

type UserAttribute struct {
	ID             string `json:"id"`
	UserID         string `json:"user_id"`
	AttributeKey   string `json:"attribute_key"`
	AttributeValue string `json:"attribute_value"`
	ValidFrom      int64  `json:"valid_from"`
	ValidTo        int64  `json:"valid_to,omitempty"`
	CreatedAt      int64  `json:"created_at"`
	UpdatedAt      int64  `json:"updated_at"`
	SourceSystem   string `json:"source_system"`
	Description    string `json:"description"`

	// 关联关系
	//User User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"user,omitempty"`
}
