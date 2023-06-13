package bootstrap

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "bedev",
	Short: "Commander what a developre wan't to do",
	Long: `
		
	`,
	TraverseChildren: true,
}

// App root of the application
type App struct {
	*cobra.Command
}

// NewApp creates new root command
func NewApp() App {
	cmd := App{
		Command: rootCmd,
	}
	// cmd.AddCommand(console.GetSubCommands(CommonModules)...)
	return cmd
}

var RootApp = NewApp()
