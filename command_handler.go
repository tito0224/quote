package quote

import (
	"errors"
	"strings"
)

type CommandHandler struct {
	Application *Application
	Commands    map[string]Command
}

func NewCommandHandler(app *Application) *CommandHandler {
	return &CommandHandler{
		Application: app,
		Commands:    map[string]Command{},
	}
}

func (handler *CommandHandler) RegisterCommand(command Command) *CommandHandler {
	if _, found := handler.Commands[command.Name()]; found {
		panic("Command with name already exists")
	}
	handler.Commands[strings.ToUpper(command.Name())] = command

	for _, alias := range command.Aliases() {
		if _, found := handler.Commands[strings.ToUpper(alias)]; found {
			panic("Alias name conflicts")
		}
		handler.Commands[alias] = command
	}
	return handler
}

func (handler *CommandHandler) ParseAndExecute(query string) error {
	parts := strings.Split(query, " ")
	if len(parts) > 0 {
		// first part should match a command name or alias
		var name = strings.ToUpper(parts[0])

		if command, ok := handler.Commands[name]; ok {
			// found command, lets execute it
			var args []string
			if len(parts) > 1 {
				args = parts[1:]
			}
			return command.Execute(handler.Application, args)
		}
	}
	return errors.New("command not found")
}
