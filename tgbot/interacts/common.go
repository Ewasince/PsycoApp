package interacts

import tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"

// commands
type BotCommand string

// keyboard
type BotButton struct {
	ButtonTitle   string
	ButtonHandler func() error
}

type ButtonsRow []BotButton
type BotKeyboard struct {
	Keyboard []ButtonsRow
}

func (b *BotKeyboard) GetKeyBoard() tg.ReplyKeyboardMarkup {
	var buttonsArray [][]tg.KeyboardButton

	for _, row := range b.Keyboard {
		var buttonsRow []tg.KeyboardButton
		for _, button := range row {
			buttonsRow = append(buttonsRow, tg.KeyboardButton{
				Text: button.ButtonTitle,
			})
		}
		buttonsArray = append(buttonsArray, buttonsRow)
	}

	keyboard := tg.ReplyKeyboardMarkup{
		Keyboard: buttonsArray,
	}
	return keyboard
}
func (b *BotKeyboard) ProcessMessage(message string) error {
	for _, row := range b.Keyboard {
		for _, button := range row {
			if button.ButtonTitle == message {
				return button.ButtonHandler()
			}
		}
	}
	return nil
}
