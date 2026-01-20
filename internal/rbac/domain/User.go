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
