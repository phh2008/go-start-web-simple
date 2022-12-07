package main

import (
	"com.gientech/equipment-data-collection/pkg/config"
	"com.gientech/equipment-data-collection/pkg/global"
	"com.gientech/equipment-data-collection/pkg/logger"
	"com.gientech/equipment-data-collection/web/middleware"
	"com.gientech/equipment-data-collection/web/router"
	"flag"
	"github.com/gin-gonic/gin"
)

func main() {
	// 命令行参数
	var configFolder string
	flag.StringVar(&configFolder, "config", "./config", "指定配置文件目录，e.g. -config ./config")
	flag.Parse()
	if configFolder == "" {
		configFolder = "./config"
	}
	// wire
	appCtx := BuildApp(config.ConfigFolder(configFolder))
	if err := appCtx.Config.Viper.Unmarshal(&global.Profile); err != nil {
		panic(err)
	}
	logger.S().Infof("----------------- start -----------------")
	// gin
	app := gin.New()
	app.Use(middleware.GinLogger)
	app.Use(middleware.GinRecovery(true))
	app.Use(middleware.Translations())
	app.Use(middleware.Cors(global.Profile.Cors))
	router.Register(app, appCtx)
	_ = app.Run(":8088")
}
