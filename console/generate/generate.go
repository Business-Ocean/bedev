package generate

import (
	"fmt"
	"os"

	"github.com/business-ocean/bedev/cmd"
	"github.com/business-ocean/bedev/common"
	"github.com/business-ocean/bedev/common/itemselect"
	"github.com/business-ocean/bedev/console/generate/gendata"
	"github.com/business-ocean/bedev/console/generate/placeholder"
	"github.com/business-ocean/bedev/console/generate/testfolder"

	"github.com/business-ocean/bedev/console/generate/uuid"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
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
	cmd.AddCommand(GetSubCommands(Modules)...)
}

func (r *generate) Run() cmd.CommandRunner {
	return func() {
		items := []list.Item{}
		for key := range cmds {
			items = append(items, itemselect.ListItem(key))
		}
		m := itemselect.NewListSelect(items, func(choice string) {
			command, ok := cmds[choice]
			if ok {
				common.ExecuteCmd(command, fx.Options())
			}
		})
		if _, err := tea.NewProgram(m).Run(); err != nil {
			fmt.Println("Error running program:", err)
			os.Exit(1)
		}
	}
}

func NewGenerateCommand() *generate {
	g := &generate{}

	return g
}
