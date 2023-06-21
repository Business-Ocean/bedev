package genid

import (
	"math/rand"
	"os"
	"sync"
	"time"

	"github.com/business-ocean/bedev/cmd"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

type uniqueId struct {
	lenght int
	count  int
}

func (r *uniqueId) Short() string {
	return "generate a random command"
}

func (r *uniqueId) Setup(cmd *cobra.Command) {
	cmd.Flags().IntVarP(&r.lenght, "lenght", "l", 10, "length of random number to generate")
	cmd.Flags().IntVarP(&r.count, "count", "c", 10, "count of random number to generated")
}

func (r *uniqueId) Run() cmd.CommandRunner {
	return func(c *color.Color) {
		var wg sync.WaitGroup
		wg.Add(1)
		for i := 0; i < r.count; i++ {
			rand.New(rand.NewSource(time.Now().Unix()))
			var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
			b := make([]rune, r.count)
			for i := range b {
				b[i] = letters[rand.Intn(len(letters))] //nolint:gosec // for better performance
			}
			c.Printf("%v\n", string(b))
		}
		wg.Done()
		os.Exit(0)
	}
}

func NewUniqueIdCommand() *uniqueId {
	return &uniqueId{}
}
