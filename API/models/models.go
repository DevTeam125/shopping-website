package models

import (
	l "github.com/DevTeam125/shopping-website/pkg/logging"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	var err error
	DB, err = gorm.Open(sqlite.Open("app.db"), &gorm.Config{})

	if err != nil {
		l.Logging.Fatalw("Couldn't Open Gorm DB", "error", err)
	}

}
