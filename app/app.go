package app

import (
	"com.gientech/equipment-data-collection/pkg/config"
	"com.gientech/equipment-data-collection/web/controller"
	"github.com/google/wire"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var AppSet = wire.NewSet(wire.Struct(new(AppContext), "*"))

type AppContext struct {
	Config   *config.Config
	DB       *gorm.DB
	HelloApi *controller.HelloController
	Logger   *zap.Logger
}
