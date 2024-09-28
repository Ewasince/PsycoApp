package commands

import (
	. "PsychoBot/teleBotStateLib"
	"PsychoBot/tryStates/context"
	"PsychoBot/tryStates/states"
)

var StartCommand = BotCommand{
	CommandMessage: "start",
	CommandHandler: func(c BotContext) (HandlerResponse, error) {
		ctx := *c.(*context.MyBotContext)

		if !ctx.IsPatientRegistered() {
			return HandlerResponse{
				NextState:      &states.RegisterState,
				TransitionType: GoState,
			}, nil
		}

		return HandlerResponse{
			NextState:      &states.FillStoryState,
			TransitionType: GoState,
		}, nil
	},
}
