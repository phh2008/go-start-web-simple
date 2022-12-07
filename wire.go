//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"com.gientech/equipment-data-collection/app"
	"com.gientech/equipment-data-collection/pkg/config"
	"com.gientech/equipment-data-collection/pkg/logger"
	"com.gientech/equipment-data-collection/util"
	"com.gientech/equipment-data-collection/web/controller"
	"github.com/google/wire"
)

func BuildApp(configFolder config.ConfigFolder) *app.AppContext {
	wire.Build(
		config.ConfigSet,
		logger.LoggerSet,
		util.InitDB,
		//dao.DaoSet,
		//service.ServiceSet,
		controller.ControllerSet,
		app.AppSet,
	)
	return new(app.AppContext)
}
