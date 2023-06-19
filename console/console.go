package console

import (
	"github.com/business-ocean/bedev/cmd"
	"github.com/business-ocean/bedev/common"
	"github.com/business-ocean/bedev/console/generate"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

var cmds = map[string]cmd.Command{
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
