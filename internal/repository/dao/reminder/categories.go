package reminder

// NewCategoriesDao 事项分类表
func NewCategoriesDao() CategoriesDao {
	return nil
}

type CategoriesDao interface {
}

/*
字段名	数据类型	约束	默认值	示例数据
id	BIGINT	PK, AUTO_INCREMENT		1
user_id	BIGINT	FK(users.id)		1
name	VARCHAR(100)	NOT NULL		"工作"
description	TEXT		NULL	"工作任务管理"
color	VARCHAR(7)		"#3B82F6"	"#3B82F6"
icon	VARCHAR(50)		NULL	"briefcase"
sort_order	INT		0	1
is_active	BOOLEAN		TRUE	TRUE
created_at	TIMESTAMP		CURRENT_TIMESTAMP	"2024-01-15 10:00:00"
updated_at	TIMESTAMP		CURRENT_TIMESTAMP	"2024-01-15 10:00:00"

*/
