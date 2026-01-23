package reminder

// NewTasksDap 任务表
func NewTasksDap() TasksDao {
	return nil
}

type TasksDao interface {
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
