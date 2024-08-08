package commands

import (
	"fmt"

	cc "github.com/danielbukowski/twitch-chatbot/internal/command_controller"
	"github.com/gempir/go-twitch-irc/v4"
)

var Ping cc.CommandFunctionSignature = func(message *twitch.PrivateMessage, ircClient *twitch.Client) error {
	ircClient.Say(message.Channel, fmt.Sprintf("Pong! @%s", message.User.DisplayName))

	return nil
}
