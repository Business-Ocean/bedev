package common

import (
	"context"
	"log"

	"github.com/business-ocean/bedev/cmd"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

func WrapSubCommand(name string, cmd cmd.Command, opt fx.Option) *cobra.Command {
	wrappedCmd := &cobra.Command{
		Use:   name,
		Short: cmd.Short(),
		Run: func(c *cobra.Command, args []string) {
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
