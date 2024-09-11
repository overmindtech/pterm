package putils

import (
	"github.com/overmindtech/pterm"
)

// TableDataFromCSV converts CSV data into pterm.TableData.
//
// Usage:
//
//	pterm.DefaultTable.WithData(putils.TableDataFromCSV(csv)).Render()
func TableDataFromCSV(csv string) (td pterm.TableData) {
	return TableDataFromSeparatedValues(csv, ",", "\n")
}
