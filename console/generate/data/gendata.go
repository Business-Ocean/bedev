package gendata

import (
	"fmt"

	"github.com/business-ocean/bedev/cmd"
	"github.com/spf13/cobra"
)

type genData struct{}

func (gd *genData) Short() string            { return "generate data based on some params" }
func (gd *genData) Setup(cmd *cobra.Command) {}
func (gd *genData) Run() cmd.CommandRunner {

	return func() {
		fmt.Println("Hello gen Data")
	}
}

func NewGenDataCommand() *genData {

	return &genData{}
}
