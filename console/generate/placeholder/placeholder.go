package placeholder

import (
	"fmt"

	"github.com/business-ocean/bedev/cmd"
	"github.com/spf13/cobra"
)

type placeholder struct{}

func (gd *placeholder) Short() string            { return "generate placeholder based on some params" }
func (gd *placeholder) Setup(cmd *cobra.Command) {}
func (gd *placeholder) Run() cmd.CommandRunner {

	return func() {
		fmt.Println("Hello placeholder")
	}
}
func NewPlaceHolderCommand() *placeholder {

	return &placeholder{}
}
