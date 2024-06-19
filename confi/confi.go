package confi

import (
	"github.com/yogi270896/hs-utils/envs"
)

type AppConfig struct {
	Server ServerConfig
}

type ServerConfig struct {
	Port            int
	APINAME         string
	LOG_FILENAME    string
	ENABLE_FILELOG  string
	PROFILE         string
	HAILSHIP_USER   string
	HAILSHIP_SECRET string
	USERURL         string
	ABC             string
}

func NewConfig() *AppConfig {
	return &AppConfig{
		Server: ServerConfig{
			ABC:             envs.GetEnv("ass", "120"),
			Port:            envs.GetEnvAsInt("PORT", 8080),
			APINAME:         envs.GetEnv("APINAME", "h-api-gatway"),
			LOG_FILENAME:    envs.GetEnv("LOG_FILENAME", "logfile.txt"),
			PROFILE:         envs.GetEnv("PROFILE", "dev"),
			ENABLE_FILELOG:  envs.GetEnv("ENABLE_FILELOG", "false"),
			HAILSHIP_USER:   envs.GetEnv("HAILSHIP_USER", "internal"),
			HAILSHIP_SECRET: envs.GetEnv("HAILSHIP_SECRET", "internal"),
			USERURL:         envs.GetMerchantService(),
		},
	}
}
