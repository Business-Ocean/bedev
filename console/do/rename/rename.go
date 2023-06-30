package rename

import (
	"fmt"
	"os"

	"github.com/business-ocean/bedev/cmd"
	"github.com/spf13/cobra"
)

type rename struct{}

func (gd *rename) Short() string            { return "generate test folder based on some params" }
func (gd *rename) Setup(cmd *cobra.Command) {}
func (gd *rename) Run() cmd.CommandRunner {

	return func() {
		fmt.Println("Hello do sample ok")
		os.Exit(0)
	}
}

func NewRenameCommand() *rename {
	return &rename{}
}
