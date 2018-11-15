package quote

type CommandExit struct{}

func (c CommandExit) Name() string {
	return "EXIT"
}

func (c CommandExit) Aliases() []string {
	return []string{"QUIT"}
}

func (c CommandExit) Description() string {
	return "Exit Quote"
}

func (c CommandExit) Execute(app *Application, args []string) error {
	app.Stop()
	return nil
}
