package configs

import (
	"go-auth-service/app/models"
	"os"
)

type AppConfigs struct {
	Database models.DBConnection
}

func NewAppConfigs() *AppConfigs {
	configs := &AppConfigs{
		Database: models.DBConnection{
			User:   os.Getenv("DBUser"),
			Passwd: os.Getenv("DBPass"),
			Addr:   os.Getenv("DBAddr"),
			Name:   os.Getenv("DBName"),
		},
	}
	return configs
}
