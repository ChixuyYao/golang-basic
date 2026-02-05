package ioc

import (
	"golang/internal/repository/dao"
	"golang/pkg/gormx"

	prometheus2 "github.com/prometheus/client_golang/prometheus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/prometheus"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open(viper.GetString("db.dsn")), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, // 禁用数据库外键约束生成
	})
	if err != nil {
		panic(err) //数据库启动失败,终止进程
	}

	// 关键指标参考:
	// gorm_dbstats_wait_count, gorm_dbstats_wait_duration: 值较大->连接数量不够,增大配置
	// gorm_dbstats_idle: 值较大->调小空闲链接数量
	// gorm_dbstats_max_idletime_closed: 值很大->调大最大空闲时间
	err = db.Use(prometheus.New(prometheus.Config{
		DBName:          "golang_basic",
		RefreshInterval: 15, // 每15s拉取数据一次
		MetricsCollector: []prometheus.MetricsCollector{
			&prometheus.MySQL{
				VariableNames: []string{"thread_running"},
			},
		},
	}))
	if err != nil {
		panic("Prometheus Gorm init fail")
	}
	cb := gormx.NewCallbacks(prometheus2.SummaryOpts{
		Namespace: "golang_basic",
		Subsystem: "user",
		Name:      "gorm_db",
		Help:      "统计 GORM 服务的数据库操作",
		ConstLabels: map[string]string{
			"instance_id": "my_instance",
		},
		Objectives: map[float64]float64{
			0.5:   0.01,
			0.75:  0.01,
			0.9:   0.005,
			0.98:  0.002,
			0.99:  0.001,
			0.999: 0.0001,
		},
	})

	err = db.Use(cb)
	if err != nil {
		panic("Prometheus Gorm Plugin init fail")
	}

	//db.Use(
	//	tracing.NewPlugin(
	//		tracing.WithoutMetrics(),
	//		tracing.WithDBSystem("MySQL_Golang"),
	//	),
	//)

	err = dao.InitTables(db)
	if err != nil {
		panic(err) //数据库表创建失败,终止进程
	}
	return db
}
