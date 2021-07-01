package databases

import (
	barroth_config "github.com/aofiee/barroth/config"
	"gorm.io/gorm"
)

var (
	DB      *gorm.DB
	dbError error
)

type (
	DBConfig struct {
		config barroth_config.Config
	}
)

func NewConfig(config barroth_config.Config) *DBConfig {
	return &DBConfig{
		config: config,
	}
}

func (db *DBConfig) DBConnString() string {
	dns := db.config.DbUser + ":" + db.config.DbPassword + "@tcp(" + db.config.DbHost + ":" + db.config.DbPort + ")/" + db.config.DbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	return dns
}
