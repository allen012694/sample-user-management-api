package config

const (
	MIGRATION_FOLDER = "migrations"
	RUNNING_PORT     = "7000"

	// TODO: move to .env
	DSN    = "root@tcp(127.0.0.1:3306)/user_db?charset=utf8mb4&parseTime=True&loc=Local"
	SALT   = "1ec8a9746ac92e99ba82"
	SECRET = "d85e60448b904daf411b"
)
