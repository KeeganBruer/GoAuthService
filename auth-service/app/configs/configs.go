package configs

import (
	"go-auth-service/app/models"
	"os"
	"strconv"
	"strings"
)

type Ports struct {
	PublicPort  int
	PrivatePort int
}
type AppConfigs struct {
	IsDev         bool
	Database      models.DBConnection
	InternalPorts Ports
	ExternalPorts Ports
}

func NewAppConfigs() *AppConfigs {
	ports := strings.Split(os.Getenv("ExternalPorts"), ",")
	internalPorts := Ports{
		PublicPort:  8080,
		PrivatePort: 8081,
	}
	externalPorts := Ports{
		PublicPort:  internalPorts.PublicPort,
		PrivatePort: internalPorts.PrivatePort,
	}
	if len(ports) > 1 {
		externalPorts.PublicPort, _ = strconv.Atoi(ports[0])
		externalPorts.PrivatePort, _ = strconv.Atoi(ports[1])
	}
	configs := &AppConfigs{
		IsDev: false,
		Database: models.DBConnection{
			User:   os.Getenv("DBUser"),
			Passwd: os.Getenv("DBPass"),
			Addr:   os.Getenv("DBAddr"),
			Name:   os.Getenv("DBName"),
		},
		InternalPorts: internalPorts,
		ExternalPorts: externalPorts,
	}
	return configs
}
