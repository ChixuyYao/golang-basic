package reminder

// NewTaskPrerequisitesDao 前置条件表
func NewTaskPrerequisitesDao() TaskPrerequisitesDao {
	return nil
}

type TaskPrerequisitesDao interface {
}

/*
字段名	数据类型	约束	默认值	示例数据
id	BIGINT	PK, AUTO_INCREMENT		1
task_id	BIGINT	FK(tasks.id)		2
prerequisite_task_id	BIGINT	FK(tasks.id)		1
condition_type	ENUM		"must_complete"	"must_complete"
required_value	INT		100	100
condition_description	VARCHAR(100)		NULL	"必须先完成数据收集"
is_satisfied	BOOLEAN		FALSE	FALSE
satisfied_at	TIMESTAMP		NULL	NULL
created_at	TIMESTAMP		CURRENT_TIMESTAMP	"2024-01-11 10:30:00"
updated_at	TIMESTAMP		CURRENT_TIMESTAMP	"2024-01-11 10:30:00"
*/

/* k扩充方案设计(v2)
字段名	数据类型	约束	默认值	描述
id	BIGINT	PK, AUTO_INCREMENT		前置条件ID
task_id	BIGINT	FK(tasks.id), NOT NULL		目标任务ID
condition_mode	ENUM('single','group','time','custom','manual')		'single'	条件模式
condition_type	ENUM('must_start','must_complete','must_progress','must_reach_checkpoint','time_based','expression_based','manual_confirmation')		'must_complete'	条件类型
prerequisite_type	ENUM('task','goal','category','system','external')		'task'	前置对象类型
prerequisite_task_id	BIGINT	FK(tasks.id)	NULL	前置任务ID
prerequisite_goal_id	BIGINT	FK(goals.id)	NULL	前置目标ID
prerequisite_category_id	BIGINT	FK(categories.id)	NULL	前置分类ID
prerequisite_config	JSON		NULL	前置条件配置
condition_expression	VARCHAR(500)		NULL	条件表达式
condition_parameters	JSON		NULL	条件参数
required_value	INT		100	要求值
operator	ENUM('and','or','xor','not')		'and'	条件运算符
group_id	BIGINT		NULL	条件组ID
parent_condition_id	BIGINT	FK(task_prerequisites.id)	NULL	父条件ID
priority	TINYINT		1	条件优先级
is_required	BOOLEAN		TRUE	是否必须
can_be_overridden	BOOLEAN		FALSE	是否可覆盖
override_reason	VARCHAR(255)		NULL	覆盖原因
override_by	BIGINT	FK(users.id)	NULL	覆盖人
override_at	TIMESTAMP		NULL	覆盖时间
is_satisfied	BOOLEAN		FALSE	是否满足
satisfied_at	TIMESTAMP		NULL	满足时间
evaluated_at	TIMESTAMP		NULL	最后评估时间
evaluation_count	INT		0	评估次数
created_at	TIMESTAMP		CURRENT_TIMESTAMP	创建时间
updated_at	TIMESTAMP	ON UPDATE CURRENT_TIMESTAMP	CURRENT_TIMESTAMP	更新时间
expired_at	TIMESTAMP		NULL	过期时间
is_active	BOOLEAN		TRUE	是否激活
*/

/* 扩充方案(v2) 字段说明
condition_mode: 条件模式
	single	单一条件	简单的前置条件
	group	条件组	多个条件的组合逻辑
	time	时间条件	基于时间的触发条件
	custom	自定义条件	复杂的业务逻辑
	manual	手动确认	需要人工确认的条件

condition_type: 条件类型说明
	must_start	必须开始	{"started": true}
	must_complete	必须完成	{"completed": true}
	must_progress	必须达到进度	{"progress": 50}
	must_reach_checkpoint	必须达到检查点	{"checkpoint_id": 1}
	time_based	时间条件	{"after_date": "2024-01-01", "before_date": "2024-12-31"}
	expression_based	表达式条件	{"expression": "progress > 50 AND status = 'in_progress'"}
	manual_confirmation	手动确认	{"confirmer_id": 1, "required": true}

prerequisite_type: 前置对象类型
	task	任务依赖	前置任务必须完成
	goal	目标依赖	前置目标必须达到特定进度
	category	分类依赖	分类下任务完成数量
	system	系统条件	用户权限、时间等
	external	外部条件	API调用、文件存在等

===> 数据库 数据示例
id	task_id	condition_mode	    condition_type	   prerequisite_type	prerequisite_task_id	prerequisite_config																	示例值
1	2	        single       	must_complete	      task	                   1	                    NULL																	任务2依赖任务1完成
2	3	        time        	time_based	          system	              NULL	            {"after_date": "2024-02-01", "before_date": "2024-02-28"}						时间窗口条件
3	4        	group       	must_progress	      task    	               5	            {"progress": 50, "operator": ">="}												进度条件组
4	5	        custom	        expression_based  	  task	                  NULL	            {"expression": "${task1.progress} > 50 AND ${task2.status} = 'completed'"}		表达式条件
5	6	        manual       	manual_confirmation	  system	              NULL	            {"required_confirmer_ids": [1,2], "minimum_confirmations": 2}					多人确认条件
*/
