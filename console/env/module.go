package env

import (
	"github.com/business-ocean/bedev/module"
	"go.uber.org/fx"
)

var envModules = fx.Options(fx.Provide(module.NewColorConsole))
