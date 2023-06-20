package generate

import (
	"github.com/business-ocean/bedev/module"
	"go.uber.org/fx"
)

var generateModules = fx.Options(fx.Provide(module.NewColorConsole))
