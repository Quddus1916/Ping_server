package config

import (
	"github.com/ping/utils/consts"
	"github.com/spf13/viper"
	"time"
)

// App holds the app configuration
type App struct {
	Name         string
	Debug        bool
	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
	CacheType    consts.CacheType
	LoggerType   consts.LoggerType
	QueueType    consts.QueueType
}

var app = &App{}

func AppCfg() *App {
	return app
}

func LoadAppCfg() {
	app.Debug = viper.GetBool("app.debug")
	app.CacheType = consts.CacheType(viper.GetString("app.cache_type"))
	app.LoggerType = consts.LoggerType(viper.GetString("app.logger_type"))
	app.QueueType = consts.QueueType(viper.GetString("app.queue_type"))
	app.HTTPPort = viper.GetInt("app.http_port")
	app.ReadTimeout = viper.GetDuration("app.read_timeout") * time.Second
	app.WriteTimeout = viper.GetDuration("app.write_timeout") * time.Second
	app.IdleTimeout = viper.GetDuration("app.idle_timeout") * time.Second
	app.Name = viper.GetString("app.name")
}
