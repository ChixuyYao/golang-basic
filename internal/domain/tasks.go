package domain

import "time"

// Tasks 领域对象:任务
type Tasks struct {
	Id               string    `json:"id"`
	GoalId           string    `json:"goal_id"`
	CategoryId       string    `json:"category_id"`
	ParentTaskId     string    `json:"parent_task_id"`
	Title            string    `json:"title"`
	Description      string    `json:"description"`
	Status           int       `json:"status"`
	Priority         int       `json:"priority"`
	Progress         float64   `json:"progress"`
	TaskType         int       `json:"task_type"`
	StartDate        time.Time `json:"start_date"`
	DueDate          time.Time `json:"due_date"`
	CompletedAt      time.Time `json:"completed_at"`
	EstimatedMinutes int64     `json:"estimated_minutes"`
	ActualMinutes    int64     `json:"actual_minutes"`
	SortOrder        int       `json:"sort_order"`
	Depth            int       `json:"depth"`
	IsLocked         int       `json:"is_locked"`
	LockReason       int       `json:"lock_reason"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
