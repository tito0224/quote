package main

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

func main() {
	resultsTable := tview.NewTable().
		SetSelectable(true, false).
		SetBorders(true).
		// SetFixed(0, 6).
		SetCellSimple(0, 0, "SYMBOL").
		SetCellSimple(0, 1, "OPEN").
		SetCellSimple(0, 2, "HIGH").
		SetCellSimple(0, 3, "LOW").
		SetCellSimple(0, 4, "CLOSE").
		SetCellSimple(0, 5, "VOLUME").
		SetCellSimple(1, 0, "MSFT").
		SetCellSimple(1, 1, "120.00").
		SetCellSimple(1, 2, "122.00").
		SetCellSimple(1, 3, "98.99").
		SetCellSimple(1, 4, "121.33").
		SetCellSimple(1, 5, "1234567890").
		SetCellSimple(2, 0, "MSFT").
		SetCellSimple(2, 1, "120.00").
		SetCellSimple(2, 2, "122.00").
		SetCellSimple(2, 3, "98.99").
		SetCellSimple(2, 4, "121.33").
		SetCellSimple(2, 5, "1234567890").
		SetCellSimple(3, 0, "MSFT").
		SetCellSimple(3, 1, "120.00").
		SetCellSimple(3, 2, "122.00").
		SetCellSimple(3, 3, "98.99").
		SetCellSimple(3, 4, "121.33").
		SetCellSimple(3, 5, "1234567890")

	commandBar := tview.NewInputField().
		SetPlaceholder("<GO>")

	commandBar.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEnter {
			commandBar.SetText("Enter")
			resultsTable.Select(0, 0)
		}
	})

	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(commandBar, 0, 1, true).
		AddItem(resultsTable, 0, 1, true)

	app := tview.NewApplication()
	app.SetRoot(flex, true).Run()
}
