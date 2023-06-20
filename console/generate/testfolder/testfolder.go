package testfolder

import (
	"fmt"

	"github.com/business-ocean/bedev/cmd"
	"github.com/spf13/cobra"
)

type testFolder struct{}

func (gd *testFolder) Short() string            { return "generate test folder based on some params" }
func (gd *testFolder) Setup(cmd *cobra.Command) {}
func (gd *testFolder) Run() cmd.CommandRunner {

	return func() {
		fmt.Println("Hello gen Data")
	}
}

func NewTestFolerCommand() *testFolder {

	return &testFolder{}
}
