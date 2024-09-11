package main

import (
	"github.com/overmindtech/pterm"
)

func main() {
	// Define the data for the heatmap
	data := [][]float32{
		{0.9, 0.2, -0.7, 0.4, -0.5, 0.6, -0.3, 0.8, -0.1, -1.0, 0.1, -0.8, 0.3},
		{0.2, -0.7, -0.5, -0.3, -0.1, 0.1, 0.3, 0.5, 0.9, -0.9, -0.7, -0.5, -0.3},
		{0.4, 0.4, -0.3, -1.0, 0.3, -0.2, -0.9, 0.5, -0.3, -1.0, 0.6, -0.2, -0.9},
		{0.9, -0.5, -0.1, 0.3, 1, -0.7, -0.3, 0.1, 0.7, -0.9, -0.5, 0.2, 0.6},
		{0.5, 0.6, 0.1, -0.2, -0.7, 0.8, 0.6, 0.1, -0.5, -0.7, 0.7, 0.3, 0.0},
	}

	// Define the axis labels for the heatmap
	headerData := pterm.HeatmapAxis{
		XAxis: []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m"},
		YAxis: []string{"1", "2", "3", "4", "5"},
	}

	// Print an informational message
	pterm.Info.Println("The following table has no rgb (supported by every terminal), no axis data and a legend.")
	pterm.Println()

	// Create the heatmap with the defined data and options, and render it
	pterm.DefaultHeatmap.
		WithData(data).
		WithBoxed(false).
		WithAxisData(headerData).
		WithLegend(false).
		WithColors(pterm.BgBlue, pterm.BgRed, pterm.BgGreen, pterm.BgYellow).
		WithLegend().
		Render()
}
