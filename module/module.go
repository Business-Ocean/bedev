package module

import (
	"github.com/fatih/color"
	"go.uber.org/fx"
)

var CMDModules = fx.Options(
	fx.Provide(NewColorConsole),
	// fx.Provide(listselect.NewListSelect),
)

// ---------------------------------- ---------------------------------- ----------------------------------

func NewColorConsole() *color.Color {
	p := NewColorParams()
	c := color.New(p.textColor).Add(p.style)
	return c
}
func NewColorParams() *ColorParams {
	return &ColorParams{
		textColor: color.FgGreen,
		style:     color.Bold,
		// background: null,
	}
}

type ColorParams struct {
	// fx.In
	textColor color.Attribute
	style     color.Attribute
}
