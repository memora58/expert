package global

import (
	"gateway/config"
	ut "github.com/go-playground/universal-translator"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	Trans  ut.Translator
	Config = &config.Config{}
)
