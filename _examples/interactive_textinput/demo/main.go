package main

import (
	"github.com/overmindtech/pterm"
)

func main() {
	// Create an interactive text input with single line input mode and show it
	result, _ := pterm.DefaultInteractiveTextInput.Show()

	// Print a blank line for better readability
	pterm.Println()

	// Print the user's answer with an info prefix
	pterm.Info.Printfln("You answered: %s", result)
}
