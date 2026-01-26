package dao

import (
	"context"
)

// NewGoalsDao 目标表
func NewGoalsDao() GoalsDao {
	return nil
}

type GoalsDao interface {
	Select() (Categories, error)
	Insert(ctx context.Context, entity Goals) error
	Update(ctx context.Context, id string, entity Goals) error
	Delete(ctx context.Context, id string) error
}

type Goals struct {
	Id          string  `gorm:"type:varchar(255);primaryKey;comment:目标ID(PK);"`
	CategoryId  string  `gorm:"type:varchar(255);comment:分类ID(FK);"`
	Title       string  `gorm:"type:varchar(255);comment:目标名称;"`
	Description string  `gorm:"comment:目标备注;"`
	Status      int     `gorm:"comment:目标状态(not_start:未开始,in_progress:进行中,completed:已完成,on_hold:暂停中,canceled:已取消);"`
	Priority    int     `gorm:"comment:目标优先级(low:,medium:,high:,critical:);"`
	StartDate   int64   `gorm:"comment:目标计划的开始时间;"`
	TargetDate  int64   `gorm:"comment:目标计划的完成截止时间(Deadline);"`
	CompletedAt int64   `gorm:"comment:目标实际完成时间;"`
	Progress    float64 `gorm:"comment:目标进度(0.00-100.00);"`
	Tags        string  `gorm:"comment:标签数组(JSON)"`
	Metrics     string  `gorm:"comment:度量指标(JSON):存储目标的关键指标"`
	CreatedAt   int64   `gorm:"comment:创建时间;"`
	UpdatedAt   int64   `gorm:"comment:更新时间;"`
}

/*
字段名			数据类型			约束					默认值				示例数据
id				BIGINT			PK, AUTO_INCREMENT						1
category_id		BIGINT			FK(categories.id)						1
title			VARCHAR(200)	NOT NULL								"完成季度报告"
description		TEXT			NULL									"完成Q1季度业务报告和演示"
status			ENUM								"not_started"		"in_progress"
priority		ENUM								"medium"			"high"
start_date		DATE			NULL									"2024-01-01"
target_date		DATE			NULL									"2024-03-31"
completed_at	TIMESTAMP		NULL				NULL
progress		TINYINT								0					30
created_at		TIMESTAMP							CURRENT_TIMESTAMP	"2024-01-01 09:00:00"
updated_at		TIMESTAMP							CURRENT_TIMESTAMP	"2024-01-15 14:00:00"'
*/
