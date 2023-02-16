package actions

import (
	"fmt"
	"net/http"
	"tg-bot/pkg/telegram/types"
)

func SetWebhook() {
	_, err := http.Get(
		fmt.Sprintf("%s%s/setWebhook?url=%s",
			types.TELEGRAM_URL,
			types.BOT_TOKEN,
			types.BOT_ADDRESS,
		))
	if err != nil {
		fmt.Println("error while get request")
	}
}
