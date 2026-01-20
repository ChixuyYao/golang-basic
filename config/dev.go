package config

var Config = config{
	DB: DBConfig{
		DSN: "root:123456@tcp(localhost:13316)/dbs",
	},
	Redis: RedisConfig{
		Addr: "localhost:6379",
	},
}

// LanguageConfig Languages数据库配置
var LanguageConfig = config{
	DB: DBConfig{
		DSN: "root:123456@tcp(localhost:13316)/languages",
	},
	Redis: RedisConfig{
		Addr: "localhost:6379",
	},
}
