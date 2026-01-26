package dao

import "context"

// NewTasksDap 任务表
func NewTasksDap() TasksDao {
	return nil
}

type TasksDao interface {
	Select() (Categories, error)
	Insert(ctx context.Context, entity Tasks) error
	Update(ctx context.Context, id string, entity Tasks) error
	Delete(ctx context.Context, id string) error
}

type Tasks struct {
	Id               string  `gorm:"type:varchar(255);primaryKey;comment:目标ID(PK);"`
	GoalId           string  `gorm:"type:varchar(255);comment:目标ID(FK);"`
	CategoryId       string  `gorm:"type:varchar(255);comment:分类ID(FK);"`
	ParentTaskId     string  `gorm:"type:varchar(255);comment:父级任务ID(FK);"`
	Title            string  `gorm:"type:varchar(255);comment:任务名称;"`
	Description      string  `gorm:"comment:任务备注;"`
	Status           int     `gorm:"comment:任务状态(pending:待处理,in_progress:进行中,completed:已完成,blocked:已阻塞,canceled:已取消,deferred:已延期);"`
	Priority         int     `gorm:"comment:任务优先级;"`
	Progress         float64 `gorm:"comment:任务进度(0.00-100.00);"`
	TaskType         int     `gorm:"comment:任务类型(category_task:分类下直接任务,sub_task:子任务);"`
	StartDate        int64   `gorm:"comment:目标计划开始时间;"`
	DueDate          int64   `gorm:"comment:目标计划截止时间;"`
	CompletedAt      int64   `gorm:"comment:目标实际完成时间;"`
	EstimatedMinutes int64   `gorm:"comment:预估时长;"`
	ActualMinutes    int64   `gorm:"comment:实际时长;"`
	SortOrder        int     `gorm:"comment:排序信息;"`
	Depth            int     `gorm:"comment:层级深度;"`
	IsLocked         int     `gorm:"comment:目标锁定状态;"`
	LockReason       int     `gorm:"comment:目标锁定原因;"`
	CreatedAt        int64   `gorm:"comment:创建时间;"`
	UpdatedAt        int64   `gorm:"comment:更新时间;"`
}

/*
字段名	数据类型	约束	默认值	示例数据1	示例数据2
id	BIGINT	PK, AUTO_INCREMENT		1	2
goal_id	BIGINT	FK(goals.id)	NULL	1	1
category_id	BIGINT	FK(categories.id)	NULL	NULL	NULL
parent_task_id	BIGINT	FK(tasks.id)	NULL	NULL	1
title	VARCHAR(200)	NOT NULL		"收集数据"	"整理数据"
description	TEXT		NULL	"收集各部门季度数据"	"整理和清洗数据"
status	ENUM		"pending"	"in_progress"	"pending"
priority	ENUM		"medium"	"high"	"medium"
task_type	ENUM		"goal_task"	"goal_task"	"subtask"
progress	TINYINT		0	50	0
due_date	DATE		NULL	"2024-01-20"	"2024-01-25"
start_date	DATE		NULL	"2024-01-10"	NULL
completed_at	TIMESTAMP		NULL	NULL	NULL
estimated_minutes	INT		NULL	300	120
actual_minutes	INT		0	150	0
sort_order	INT		0	1	2
depth	TINYINT		0	0	1
is_locked	BOOLEAN		FALSE	FALSE	TRUE
lock_reason	VARCHAR(255)		NULL	NULL	"等待数据收集完成"
created_at	TIMESTAMP		CURRENT_TIMESTAMP	"2024-01-10 09:00:00"	"2024-01-11 10:00:00"
updated_at	TIMESTAMP		CURRENT_TIMESTAMP	"2024-01-15 11:00:00"	"2024-01-15 11:00:00"
*/
