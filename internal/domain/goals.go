package domain

import "time"

// Goals 领域对象:目标
type Goals struct {
	Id          string
	CategoryId  string
	Title       string
	Description string
	Status      int
	Priority    int
	StartDate   time.Time
	TargetDate  time.Time
	CompletedAt time.Time
	Progress    float64
	Tags        string
	Metrics     string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
