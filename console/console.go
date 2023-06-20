package console

import (
	"github.com/business-ocean/bedev/cmd"
	"github.com/business-ocean/bedev/common"
	"github.com/business-ocean/bedev/console/clean"
	"github.com/business-ocean/bedev/console/config"
	"github.com/business-ocean/bedev/console/convert"
	"github.com/business-ocean/bedev/console/create"
	"github.com/business-ocean/bedev/console/do"
	"github.com/business-ocean/bedev/console/download"
	"github.com/business-ocean/bedev/console/env"
	"github.com/business-ocean/bedev/console/generate"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

var cmds = map[string]cmd.Command{
	"clean":    clean.NewCleanCommand(),
	"config":   config.NewConfigCommand(),
	"convert":  convert.NewConvertCommand(),
	"create":   create.NewCreateCommand(),
	"do":       do.NewDoCommand(),
	"download": download.NewDownloadCommand(),
	"env":      env.NewEnvCommand(),
	"generate": generate.NewGenerateCommand(),
}

// GetSubCommands gives a list of sub commands
func GetSubCommands(opt fx.Option) []*cobra.Command {
	subCommands := make([]*cobra.Command, 0)
	for name, cmd := range cmds {
		subCommands = append(subCommands, common.WrapSubCommand(name, cmd, opt))
	}
	return subCommands
}
