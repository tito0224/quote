package quote

type Command interface {
	Name() string
	Description() string
	Aliases() []string
	Execute(app *Application, args []string) error
}
