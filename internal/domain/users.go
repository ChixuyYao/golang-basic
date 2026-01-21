package domain

import "time"

// 领域对象层
type User struct {
	Id       int64
	Email    string
	Password string

	CTime time.Time
	UTime time.Time
}

type Users struct {
	Id        string    `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	State     int       `json:"state"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	LastLogin time.Time `json:"last_login"`
	// 关联关系
	//Attributes []UserAttribute         `gorm:"foreignKey:UserID" json:"attributes,omitempty"`
	//AccessLogs []AccessLog            `gorm:"foreignKey:UserID" json:"access_logs,omitempty"`
	//Contexts   []EnvironmentContext   `gorm:"foreignKey:UserID" json:"contexts,omitempty"`
	//Policies   []Policy               `gorm:"many2many:policy_targets;" json:"policies,omitempty"`
}
