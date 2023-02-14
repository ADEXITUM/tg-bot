package telegram

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func InitBot() *tgbotapi.BotAPI {
	NewBot, err := tgbotapi.NewBotAPI("5845449188:AAGPf7Vc14jCbDo5XbjN8eNSghJ4qObWP4c")
	if err != nil {
		fmt.Println("error while initializing bot ", err)
	}
	return NewBot
}

func SendMessage(chatID int64, text string) {
	NewBot := InitBot()
	msg := tgbotapi.NewMessage(572377674, "Hello from aonther func")
	NewBot.Send(msg)
}
