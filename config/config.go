package config

func InitConfig() {
	initServerConfig()
	initDatabaseConfig()
	InitGorm()
}
