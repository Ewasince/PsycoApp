package states

import (
	img "PsychoBot/images"
	msg "PsychoBot/messages"
	. "github.com/Ewasince/go-telegram-state-bot/message_types"

	. "github.com/Ewasince/go-telegram-state-bot/states"
)

var InfoState = NewBotState(
	"Info state",
	BotMessages{img.InfoImage, TextMessage(msg.StartInfo)},
	nil,
	&InfoKeyboard,
	nil,
)
