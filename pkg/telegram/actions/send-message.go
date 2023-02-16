package actions

import (
	"fmt"
	"net/http"
	"tg-bot/pkg/telegram/types"
)

func SendMessage(chatID int64, text string) {
	_, err := http.Get(
		fmt.Sprintf("%s%s/sendMessage?chat_id=%d&text=%s",
			types.TELEGRAM_URL,
			types.BOT_TOKEN,
			chatID,
			text))
	if err != nil {
		fmt.Println("error while get request")
	}
}
