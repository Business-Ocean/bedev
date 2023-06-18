package generate

import (
	"math/rand"
	"time"

	"github.com/business-ocean/bedev/cmd"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

type randomCommand struct {
	num int
}

func (r *randomCommand) Short() string {
	return "generate a random command"
}

func (r *randomCommand) Setup(cmd *cobra.Command) {
	cmd.Flags().IntVarP(&r.num, "num", "n", 10, "length of random number to generate")
}

func (r *randomCommand) Run() cmd.CommandRunner {
	return func(c *color.Color) {
		// l.Info("running random command")

		rand.New(rand.NewSource(time.Now().Unix()))
		var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
		b := make([]rune, r.num)
		for i := range b {
			b[i] = letters[rand.Intn(len(letters))] //nolint:gosec // for better performance
		}
		// fmt.Println(string(b))
		// log.Info(string(b))

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

func NewUUIDCommand() *randomCommand {
	return &randomCommand{}
}
