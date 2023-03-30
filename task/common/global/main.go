package global

import (
	"gorm.io/gorm"
	"task/config"
)

var (
	DB     *gorm.DB
	Config = &config.Config{}
)
