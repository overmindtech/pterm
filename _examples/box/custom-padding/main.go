package main

import "github.com/overmindtech/pterm"

func main() {
	// Create a default box with custom padding options and print "Hello, World!" inside it.
	pterm.DefaultBox.WithRightPadding(10).WithLeftPadding(10).WithTopPadding(2).WithBottomPadding(2).Println("Hello, World!")
}
