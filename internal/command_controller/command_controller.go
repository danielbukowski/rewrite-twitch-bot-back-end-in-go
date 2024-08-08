package commandcontroller

import (
	"github.com/gempir/go-twitch-irc/v4"
)

const COMMAND_PREFIX string = "!"

type CommandFunctionSignature func(privateMessage *twitch.PrivateMessage, ircClient *twitch.Client) error

type commandController struct {
	commands map[string]func(privateMessage *twitch.PrivateMessage, ircClient *twitch.Client) error
}

func NewCommandController() *commandController {
	return &commandController{
		commands: make(map[string]func(message *twitch.PrivateMessage, ircClient *twitch.Client) error),
	}
}

func (cc *commandController) CallCommand(commandName string, privateMessage *twitch.PrivateMessage, ircClient *twitch.Client) {
	command := cc.commands[commandName]
	if command == nil {
		return
	}

	command(privateMessage, ircClient)
}

func (cc *commandController) AddCommand(commandName string, command CommandFunctionSignature) {
	cc.commands[commandName] = command
}
