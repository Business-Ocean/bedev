package sample

import (
	"fmt"
	"os"

	"github.com/business-ocean/bedev/cmd"
	"github.com/spf13/cobra"
)

type sample struct{}

func (gd *sample) Short() string            { return "generate test folder based on some params" }
func (gd *sample) Setup(cmd *cobra.Command) {}
func (gd *sample) Run() cmd.CommandRunner {

	return func() {
		fmt.Println("Hello create sample")
		os.Exit(0)
	}
}

func NewSampleCommand() *sample {

	return &sample{}
}
