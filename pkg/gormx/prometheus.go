package gormx

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"gorm.io/gorm"
)

type Callbacks struct {
	vector *prometheus.SummaryVec
}

func NewCallbacks(opt prometheus.SummaryOpts) *Callbacks {
	// 期望监控数据:操作类型(type),操作的表(table)
	vector := prometheus.NewSummaryVec(opt, []string{"type", "table"})
	prometheus.MustRegister(vector)

	return &Callbacks{
		vector: vector,
	}
}

func (c *Callbacks) Before() func(db *gorm.DB) {
	return func(db *gorm.DB) {
		// SQL 开始时间
		start := time.Now()
		db.Set("start_time", start)
	}
}

func (c *Callbacks) After(typ string) func(db *gorm.DB) {
	return func(db *gorm.DB) {
		// SQL 结束时间
		val, _ := db.Get("start_time")
		start, ok := val.(time.Time)
		if ok {
			duration := time.Since(start).Milliseconds()
			c.vector.WithLabelValues(typ, db.Statement.Table).
				Observe(float64(duration))
		}
	}
}

// Name GORM 插件名
func (c *Callbacks) Name() string {
	return "Prometheus"
}

// Initialize GORM 插件
func (c *Callbacks) Initialize(db *gorm.DB) error {
	err := db.Callback().Create().Before("*").Register("prometheus_gorm_create_before", c.Before())
	if err != nil {
		return err
	}
	err = db.Callback().Create().After("*").Register("prometheus_gorm_create_after", c.After("CREATE"))
	if err != nil {
		return err
	}
	err = db.Callback().Query().Before("*").Register("prometheus_gorm_query_before", c.Before())
	if err != nil {
		return err
	}
	err = db.Callback().Query().After("*").Register("prometheus_gorm_query_after", c.After("QUERY"))
	if err != nil {
		return err
	}
	err = db.Callback().Raw().Before("*").Register("prometheus_gorm_raw_before", c.Before())
	if err != nil {
		return err
	}
	err = db.Callback().Raw().After("*").Register("prometheus_gorm_raw_after", c.After("RAW"))
	if err != nil {
		return err
	}
	err = db.Callback().Update().Before("*").Register("prometheus_gorm_update_before", c.Before())
	if err != nil {
		return err
	}
	err = db.Callback().Update().After("*").Register("prometheus_gorm_update_after", c.After("UPDATE"))
	if err != nil {
		return err
	}
	err = db.Callback().Delete().Before("*").Register("prometheus_gorm_delete_before", c.Before())
	if err != nil {
		return err
	}
	err = db.Callback().Delete().After("*").Register("prometheus_gorm_delete_after", c.After("DELETE"))
	if err != nil {
		return err
	}
	err = db.Callback().Row().Before("*").Register("prometheus_gorm_row_before", c.Before())
	if err != nil {
		return err
	}
	err = db.Callback().Row().After("*").Register("prometheus_gorm_row_after", c.After("ROW"))
	return err
}
