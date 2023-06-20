package systemenv

import (
	"os"
	"strings"

	"github.com/business-ocean/bedev/cmd"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

type systemenv struct{}

func (gd *systemenv) Short() string            { return "generate test folder based on some params" }
func (gd *systemenv) Setup(cmd *cobra.Command) {}
func (gd *systemenv) Run() cmd.CommandRunner {

	return func(c *color.Color) {

		c.Printf("\n%v\n", "Fetching you env variable")

		for _, e := range os.Environ() {
			pair := strings.SplitN(e, "=", -1)

			key := pair[0]

			if key == "PATH" {
				c.Print(key)
				value := pair[1]
				paths := strings.Split(value, ":")
				for _, path := range paths {
					c.Println("	 " + path)

				}
				c.Println()
				continue
			}

			c.Add(color.FgBlue)
			c.Print(key)
			c.Print(" = ")
			// c.Add(color.BgGreen)
			color.Magenta(pair[1] + "\n")

		}

		os.Exit(0)
	}
}

func NewSampleCommand() *systemenv {

	return &systemenv{}
}
