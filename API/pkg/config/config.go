package config

import (
	l "github.com/DevTeam125/shopping-website/pkg/logging"
	"github.com/go-ini/ini"
)

var cfg *ini.File

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

var DatabaseSetting = &Database{}

func Init() {
	var err error
	cfg, err = ini.Load("config/app.ini")
	if err != nil {
		l.Logging.Fatalw("Failed to parse 'app.ini'", "error", err)
	}

	mapTo("database", DatabaseSetting)
}
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		l.Logging.Fatalw("Failed to map", "error", err)
	}
}
