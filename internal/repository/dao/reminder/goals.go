package reminder

// NewGoalsDao 目标表
func NewGoalsDao() GoalsDao {
	return nil
}

type GoalsDao interface {
}

/*
字段名	数据类型	约束	默认值	示例数据
id	BIGINT	PK, AUTO_INCREMENT		1
category_id	BIGINT	FK(categories.id)		1
title	VARCHAR(200)	NOT NULL		"完成季度报告"
description	TEXT		NULL	"完成Q1季度业务报告和演示"
status	ENUM		"not_started"	"in_progress"
priority	ENUM		"medium"	"high"
start_date	DATE		NULL	"2024-01-01"
target_date	DATE		NULL	"2024-03-31"
completed_at	TIMESTAMP		NULL	NULL
progress	TINYINT		0	30
created_at	TIMESTAMP		CURRENT_TIMESTAMP	"2024-01-01 09:00:00"
updated_at	TIMESTAMP		CURRENT_TIMESTAMP	"2024-01-15 14:00:00"'
*/
