package generate

import (
	"fmt"
	"os"

	"github.com/business-ocean/bedev/cmd"
	"github.com/business-ocean/bedev/common"
	listselect "github.com/business-ocean/bedev/common/list_select"
	gendata "github.com/business-ocean/bedev/console/generate/data"
	"github.com/business-ocean/bedev/console/generate/placeholder"
	testfolder "github.com/business-ocean/bedev/console/generate/test_folder"
	"github.com/business-ocean/bedev/console/generate/uuid"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

var cmds = map[string]cmd.Command{
	"uuid":        uuid.NewUUIDCommand(),
	"gendata":     gendata.NewGenDataCommand(),
	"placeholder": placeholder.NewPlaceHolderCommand(),
	"test_folder": testfolder.NewTestFolerCommand(),
}

// GetSubCommands gives a list of sub commands
func GetSubCommands(opt fx.Option) []*cobra.Command {
	subCommands := make([]*cobra.Command, 0)
	for name, cmd := range cmds {
		subCommands = append(subCommands, common.WrapSubCommand(name, cmd, opt))
	}
	return subCommands
}

type generate struct {
}

func (r *generate) Short() string {
	return "Generate command used for generate some cool developer stuffs."
}

func (r *generate) Setup(cmd *cobra.Command) {
	// cmd.Flags().IntVarP(&r.num, "num", "n", 10, "length of random number to generate")
}

func (r *generate) Run() cmd.CommandRunner {
	return func(c *color.Color, m *listselect.ListModel) {

		if _, err := tea.NewProgram(m).Run(); err != nil {
			fmt.Println("Error running program:", err)
			os.Exit(1)
		}

	}
}

func NewGenerateCommand() *generate {
	return &generate{}
}
