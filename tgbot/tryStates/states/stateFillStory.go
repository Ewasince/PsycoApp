package states

import (
	msg "PsychoBot/messages"
	. "PsychoBot/teleBotStateLib"
	"PsychoBot/tryStates"
	"PsychoBot/tryStates/helpers"
	"StorageModule/repo"
	"strconv"
)

var FillStoryState = newBotStateWrapper(
	"Fill Schedule state",
	BotMessageHandler(enterMessageHandlerFillStoryState),
	nil,
	nil,
	messageHandlerFillStoryState,
)

func enterMessageHandlerFillStoryState(c BotContext) ([]string, error) {
	ctx := *c.(*tryStates.MyBotContext)
	ctx.NewStory()
	return []string{msg.WhatHappened}, nil
}
func messageHandlerFillStoryState(c BotContext) (HandlerResponse, error) {
	ctx := *c.(*tryStates.MyBotContext)

	story := ctx.GetStory()

	if story.Situation == "" {
		story.Situation = ctx.MessageText
		_ = ctx.CreateAndSendMessage(msg.WhatMind)
		return HandlerResponse{}, nil
	} else if story.Mind == "" {
		story.Mind = ctx.MessageText
		_ = ctx.CreateAndSendMessage(msg.WhatEmotion)
		return HandlerResponse{}, nil
	} else if story.Emotion == "" {
		story.Emotion = ctx.MessageText
		_ = ctx.CreateAndSendMessage(msg.WhatPower)
		return HandlerResponse{}, nil
	} else if story.Power == 0 {
		power, err := strconv.Atoi(ctx.MessageText)
		if err != nil {
			_ = ctx.CreateAndSendMessage(msg.DontRecognizePower)
			return HandlerResponse{}, nil
		}
		if power < 1 || power > 10 {
			_ = ctx.CreateAndSendMessage(msg.DontRecognizePower)
			return HandlerResponse{}, nil
		}
		story.Power = uint8(power)
	}

	story.Date = helpers.GetDate()
	err := repo.CreateStory(story)

	if err != nil {
		_ = ctx.CreateAndSendMessage(msg.CantSaveStory)
		return HandlerResponse{}, nil
	}
	_ = ctx.CreateAndSendMessage(msg.WhatEntryDone)
	_ = ctx.NewStory()
	return HandlerResponse{
		NextState:      nil,
		TransitionType: ReloadState,
	}, nil
}