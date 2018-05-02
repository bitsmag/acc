package useroutput

import (
	"acc/types"

	"github.com/fatih/color"
)

// PrintPrintablelns prints one or more stirings in separate lines with the provided textcolor
func PrintPrintablelns(printables ...types.Printable) {
	for _, p := range printables {
		var c *color.Color
		switch p.Color() {
		case "green":
			c = color.New(color.FgGreen)
		case "red":
			c = color.New(color.FgRed)
		default:
			c = color.New(color.FgWhite)
		}
		c.Println(p.String())
	}
}
