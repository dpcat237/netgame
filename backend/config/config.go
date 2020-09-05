package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

// ModeProd defines production mode
const ModeProd = Mode("prod")

type Mode string

// Config defines configuration parameters
type Config struct {
	DbDSN         string `envconfig:"DB_DSN" default:"root:root@tcp(db_container:3306)/backend?charset=utf8mb4"`
	HTTPport      string `envconfig:"HTTP_PORT" default:"8081"`
	WebsocketPort string `envconfig:"WEBSOCKET_PORT" default:"8085"`
	Mode          Mode   `envconfig:"MODE" default:"dev"`
}

// LoadConfigData loads environment parameters
func LoadData() Config {
	var cnf Config
	if err := envconfig.Process("boilerplate", &cnf); err != nil {
		panic(fmt.Sprintf("Failed reading environment variables: %s", err))
	}
	return cnf
}

func (md Mode) IsProduction() bool {
	return md == ModeProd
}
