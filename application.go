package quote

import (
	"unicode"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

type Application struct {
	app       *tview.Application
	pages     *tview.Pages
	statusBox *tview.TextView
}

func NewApplication() *Application {
	app := tview.NewApplication()
	pages := tview.NewPages()
	messageBox := tview.NewTextView()

	quoteApp := &Application{
		app:       app,
		pages:     pages,
		statusBox: messageBox,
	}

	commandBar := tview.NewInputField().
		SetPlaceholder("<GO>")

	grid := tview.NewGrid().
		SetRows(1, 1, -1).
		SetColumns(0).
		SetBorders(true).
		AddItem(commandBar, 0, 0, 1, 1, 0, 0, true).
		AddItem(messageBox, 1, 0, 1, 1, 0, 0, false).
		AddItem(pages, 2, 0, 1, 1, 0, 0, false)

	commandHandler := NewCommandHandler(quoteApp).
		RegisterCommand(CommandExit{}).
		RegisterCommand(CommandQuote{})

	commandBar.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEnter {
			err := commandHandler.ParseAndExecute(commandBar.GetText())
			if err != nil {
				quoteApp.ShowError(err.Error())
			}
		}
	})

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		// FORCE ALL INPUT TO BE UPPERCASE
		if event.Key() == tcell.KeyRune {
			return tcell.NewEventKey(event.Key(), unicode.ToUpper(event.Rune()), event.Modifiers())
		}

		if event.Key() == tcell.KeyEsc ||
			event.Key() == tcell.KeyESC ||
			event.Key() == tcell.KeyEscape {
			messageBox.Clear()
			commandBar.SetText("")
		}
		return event
	})

	app.SetRoot(grid, true)

	return quoteApp
}

func (app *Application) Run() {
	app.app.Run()
}

func (app *Application) Stop() {
	app.app.Stop()
}

func (app *Application) GetPages() *tview.Pages {
	return app.pages
}

func (app *Application) ShowMessage(message string) {
	app.statusBox.SetText(message).SetTextColor(tcell.ColorDefault)
}

func (app *Application) ShowError(message string) {
	app.statusBox.SetText(message).SetTextColor(tcell.ColorRed)
}
