package console

import (
	"context"
	"log"

	"github.com/business-ocean/bedev/cmd"
	"github.com/business-ocean/bedev/console/generate"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

var cmds = map[string]cmd.Command{
	"uuid": generate.NewUUIDCommand(),
}

// GetSubCommands gives a list of sub commands
func GetSubCommands(opt fx.Option) []*cobra.Command {
	subCommands := make([]*cobra.Command, 0)
	for name, cmd := range cmds {
		subCommands = append(subCommands, WrapSubCommand(name, cmd, opt))
	}
	return subCommands
}

func WrapSubCommand(name string, cmd cmd.Command, opt fx.Option) *cobra.Command {
	wrappedCmd := &cobra.Command{
		Use:   name,
		Short: cmd.Short(),
		Run: func(c *cobra.Command, args []string) {
			cmd.Run()
			opts := fx.Options(
				fx.Invoke(cmd.Run()),
			)
			ctx := context.Background()
			app := fx.New(
				opt,
				opts,
				fx.NopLogger, // disable fx log message
			)
			err := app.Start(ctx)
			defer func() {
				err = app.Stop(ctx)
			}()
			if err != nil {
				log.Fatal(err)
			}

		},
	}
	cmd.Setup(wrappedCmd)
	return wrappedCmd
}
