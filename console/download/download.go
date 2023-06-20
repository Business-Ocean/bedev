package download

import (
	"fmt"
	"os"

	"github.com/business-ocean/bedev/cmd"
	"github.com/business-ocean/bedev/common"
	"github.com/business-ocean/bedev/common/itemselect"
	"github.com/business-ocean/bedev/console/download/sample"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

var cmds = map[string]cmd.Command{
	"sample": sample.NewSampleCommand(),
}

// GetSubCommands gives a list of sub commands
func getSubCommands(opt fx.Option) []*cobra.Command {
	subCommands := make([]*cobra.Command, 0)
	for name, cmd := range cmds {
		subCommands = append(subCommands, common.WrapSubCommand(name, cmd, opt))
	}
	return subCommands
}

type download struct {
}

func (r *download) Short() string {
	return "download command used for download some cool developer stuffs."
}

func (r *download) Setup(cmd *cobra.Command) {
	cmd.AddCommand(getSubCommands(downloadModules)...)
}

func (r *download) Run() cmd.CommandRunner {
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

func NewDownloadCommand() *download {
	g := &download{}

	return g
}
