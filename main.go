/*
Copyright © 2023 Sourav Kumar <businessocean.pvt.ltd@gmail.com>
This file is part of CLI application foo.
*/
package main

import (
	"os"

	"github.com/business-ocean/bedev/app"
	"go.uber.org/fx"
)

func main() {

	// _ = app.Bedev.Execute()

	fx.New(
		fx.NopLogger,
		fx.Options(
			fx.Provide(app.NewBedevApp),
			fx.Invoke(func(bedev *app.BedevApp) {
				err := bedev.Execute()

				if err == nil {
					os.Exit(0)
				}
				os.Exit(1)

			}),
		)).Run()
}
