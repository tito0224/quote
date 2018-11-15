package quote

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/gdamore/tcell"

	"github.com/rivo/tview"
	"github.com/tito0224/go-alpha-vantage"
)

type CommandQuote struct{}

func (c CommandQuote) Name() string {
	return "QUOTE"
}

func (c CommandQuote) Aliases() []string {
	return []string{"Q"}
}

func (c CommandQuote) Description() string {
	return "Get quote for <SYMBOL>"
}

func (c CommandQuote) Execute(app *Application, args []string) error {
	if args == nil || len(args) == 0 {
		return errors.New("You must specify at least one symbol. QUOTE <SYMBOL>")
	}

	app.ShowMessage(fmt.Sprintf("Loading quote for %s ....", strings.Join(args, ", ")))

	resultsTable := tview.NewTable().
		SetSelectable(true, false).
		SetBorders(true).
		SetCellSimple(0, 0, "SYMBOL").
		SetCellSimple(0, 1, "PRICE").
		SetCellSimple(0, 2, "CHANGE").
		SetCellSimple(0, 3, "PERCENT").
		SetCellSimple(0, 4, "OPEN").
		SetCellSimple(0, 5, "HIGH").
		SetCellSimple(0, 6, "LOW").
		SetCellSimple(0, 7, "CLOSE").
		SetCellSimple(0, 8, "VOLUME")

	client := alphago.NewDefaultClient("L4LML6FM48U076MW")

	for i, symbol := range args {
		quote, err := client.GetQuote(symbol)
		if err != nil {
			return err
		}

		change, _ := strconv.ParseFloat(quote.Change, 64)
		color := tcell.ColorDefault
		arrow := ""
		if change > 0 {
			color = tcell.ColorGreen
			arrow = "↑"
		} else if change < 0 {
			color = tcell.ColorRed
			arrow = "↓"
		}

		currentPriceCell := tview.NewTableCell(fmt.Sprintf("%s %s", arrow, quote.Price)).SetTextColor(color)
		changeCell := tview.NewTableCell(quote.Change).SetTextColor(color)
		cahngeCellPer := tview.NewTableCell(quote.ChangePercent).SetTextColor(color)

		resultsTable.
			SetCellSimple(i+1, 0, quote.Symbol).
			SetCell(i+1, 1, currentPriceCell).
			SetCell(i+1, 2, changeCell).
			SetCell(i+1, 3, cahngeCellPer).
			SetCellSimple(i+1, 4, quote.Open).
			SetCellSimple(i+1, 5, quote.High).
			SetCellSimple(i+1, 6, quote.Low).
			SetCellSimple(i+1, 7, quote.PrevClose).
			SetCellSimple(i+1, 8, quote.Volume)
	}

	app.ShowMessage(fmt.Sprintf("Loading quote for %s DONE", strings.Join(args, ", ")))
	app.GetPages().AddAndSwitchToPage("quotes", resultsTable, true)
	return nil
}
