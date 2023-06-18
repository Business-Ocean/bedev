package generate

import (
	"math/rand"
	"time"

	"github.com/business-ocean/bedev/cmd"
	"github.com/sirupsen/logrus"
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
	return func(log *logrus.Logger) {
		// l.Info("running random command")
		rand.New(rand.NewSource(time.Now().Unix()))
		var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
		b := make([]rune, r.num)
		for i := range b {
			b[i] = letters[rand.Intn(len(letters))] //nolint:gosec // for better performance
		}
		// fmt.Println(string(b))
		log.Info(string(b))
	}
}

func NewUUIDCommand() *randomCommand {
	return &randomCommand{}
}
