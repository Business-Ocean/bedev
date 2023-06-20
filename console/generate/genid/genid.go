package genid

import (
	"math/rand"
	"os"
	"time"

	"github.com/business-ocean/bedev/cmd"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

type uniqueId struct {
	num int
}

func (r *uniqueId) Short() string {
	return "generate a random command"
}

func (r *uniqueId) Setup(cmd *cobra.Command) {
	cmd.Flags().IntVarP(&r.num, "num", "n", 10, "length of random number to generate")
}

func (r *uniqueId) Run() cmd.CommandRunner {
	return func(c *color.Color) {

		rand.New(rand.NewSource(time.Now().Unix()))
		var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
		b := make([]rune, r.num)
		for i := range b {
			b[i] = letters[rand.Intn(len(letters))] //nolint:gosec // for better performance
		}

		c.Printf("Unique Id : %v\n", string(b))
		os.Exit(0)
	}
}

func NewUUIDCommand() *uniqueId {
	return &uniqueId{}
}
