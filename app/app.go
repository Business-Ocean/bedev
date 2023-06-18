package app

import (
	"github.com/business-ocean/bedev/console"
	"github.com/business-ocean/bedev/module"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "bedev",
	Short: "bedev what a developre wan't to do",
	Long: `
Bedev: The Terminal App for Developers

Bedev is a powerful terminal app that helps developers save time and get more done. With Bedev, you can easily rename files, clean project build files, and more.

Here are just a few of the things Bedev can do:

Rename files with an underscore pattern
Clean project build files like node_module
Run shell commands
Open files in a text editor
Search for files
And more!
Bedev is free to use and available for download now.

Download Bedev and start saving time today!

Here are some additional benefits of using Bedev:

Bedev is easy to use. Even if you're not a terminal expert, you can still use Bedev to get things done.
Bedev is powerful. Bedev can do a lot of things, so you can save time by using it for multiple tasks.
Bedev is free. There's no cost to use Bedev, so you can try it out without any risk.
If you're a developer, I encourage you to download Bedev and try it out. I think you'll find it to be a valuable tool that can help you save time and get more done.

`,
	TraverseChildren: true,
}

// App root of the application
type bedevApp struct {
	*cobra.Command
}

// NewApp creates new root command
func newBedevApp() bedevApp {
	cmd := bedevApp{
		Command: rootCmd,
	}
	cmd.AddCommand(console.GetSubCommands(module.CommonModules)...)
	return cmd
}

var Bedev = newBedevApp()
