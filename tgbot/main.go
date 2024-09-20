package main

import (
	. "EnvironmentModule"
	"PsychoBot/interacts"
	"PsychoBot/scheduler"
	"PsychoBot/stateBot"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func main() {
	botAPI, err := tg.NewBotAPI(Env.BOT_TOKEN)
	if err != nil {
		log.Fatal(err)
	}

	u := tg.NewUpdate(0)
	u.Timeout = 60

	updates := botAPI.GetUpdatesChan(u)

	go scheduler.Start()

	for update := range updates {
		if update.Message == nil {
			continue
		}

		messageMessage := update.Message
		messageSender := messageMessage.From
		messageCommand := messageMessage.Command()

		stateHandler := stateBot.NewStateHandler(
			messageMessage,
			botAPI,
		)

		log.Printf(
			"[%s, %d] %s",
			messageSender.UserName,
			messageSender.ID,
			update.Message.Text,
		)

		switch messageCommand {
		case interacts.StartCommandButton:
			stateBot.StatesCache.ResetState(messageSender.ID)
		}

		state := stateBot.StatesCache.GetState(messageSender.ID)
		stateHandler.ProcessState(state)
	}
}
