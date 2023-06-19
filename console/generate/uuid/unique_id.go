package uuid

import (
	"math/rand"
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

		c.Println("Prints cyan text with an underline.")

		// Or just add them to New()
		d := color.New(color.FgCyan, color.Bold)
		d.Printf("This prints bold cyan %s\n", "too!.")

		// Mix up foreground and background colors, create new mixes!
		red := color.New(color.FgRed)

		boldRed := red.Add(color.Bold)
		boldRed.Println("This will print text in bold red.")

		whiteBackground := red.Add(color.BgWhite)
		whiteBackground.Println("Red text with white background.")

	}
}

func NewUUIDCommand() *uniqueId {
	return &uniqueId{}
}
