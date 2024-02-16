package app

import "github.com/yusufcanb/tlama/pkg/config"

type BaseApp interface {
	SetConfig(config config.TlamaConfig)
	GetConfig() *config.TlamaConfig
}
