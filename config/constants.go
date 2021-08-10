package config

const (
	MIGRATION_FOLDER = "migrations"
	RUNNING_PORT     = "7000"

	// TODO: move to .env
	REDIS_ADDRESS  = "localhost:6379"
	REDIS_PASSWORD = ""
	DATABASE       = "root@tcp(127.0.0.1:3306)/user_db?charset=utf8mb4&parseTime=True&loc=Local"
	SALT           = "1ec8a9746ac92e99ba82"
	SECRET         = "4D635166546A576E5A7234753778214125442A462D4A614E645267556B587032"
)

const (
	CONTEXT_CURRENT_USER    = "current_user"
	REDIS_SESSION_STORE_KEY = "session_store"
)

const (
	LOG_ACTION_LOGIN  = "login"
	LOG_ACTION_UPDATE = "update"
	LOG_ACTION_CREATE = "create"
)
