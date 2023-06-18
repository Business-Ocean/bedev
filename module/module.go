package module

import (
	"github.com/fatih/color"
	"go.uber.org/fx"
)

var CMDModules = fx.Options(
	fx.Provide(color.New),
)
