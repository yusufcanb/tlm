package app

import "tlama/pkg/config"

type BaseApp interface {
	SetConfig(config config.TlamaConfig)
	GetConfig() *config.TlamaConfig
}
