package global

import (
	"gorm.io/gorm"
	"user/config"
)

var (
	DB     *gorm.DB
	Config = &config.Config{}
)
