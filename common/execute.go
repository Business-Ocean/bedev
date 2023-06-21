package common

import (
	"github.com/business-ocean/bedev/cmd"
	"github.com/business-ocean/bedev/module"
	"go.uber.org/fx"
)

func ExecuteCmd(cmd cmd.Command, opt fx.Option) {

	opts := fx.Options(
		fx.Provide(module.NewColorConsole),
		fx.Invoke(cmd.Run()),
	)
	//  TODO(sourav): try running fx app start and stop
	fx.New(
		opt,
		opts,
		fx.NopLogger, // disable fx log message
	).Run()

}
