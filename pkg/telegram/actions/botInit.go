package actions

import (
	"fmt"
	"net/http"
	"tg-bot/pkg/telegram/types"
)

func InitBot() {
	_, err := http.Get(
		fmt.Sprintf("%s%s/getMe",
			types.TELEGRAM_URL,
			types.BOT_TOKEN,
		))
	if err != nil {
		fmt.Println("error while get request: ", err)
	}
}
