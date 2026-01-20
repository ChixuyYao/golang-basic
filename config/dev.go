package config

var BookConfig = config{
	DB: DBConfig{
		DSN: "root:123456@tcp(localhost:13316)/books",
	},
	Redis: RedisConfig{
		Addr: "localhost:6379",
	},
}
