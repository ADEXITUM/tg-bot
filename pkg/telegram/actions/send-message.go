package actions

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"tg-bot/pkg/telegram/types"
)

func SendMessage(chatID int64, text string) {

	url := fmt.Sprintf("%s%s/sendMessage", types.TELEGRAM_URL, types.BOT_TOKEN)

	payload := strings.NewReader(fmt.Sprintf("{\"text\":\"%s\",\"parse_mode\":\"HTML\",\"disable_web_page_preview\":false,\"disable_notification\":false,\"reply_to_message_id\":null,\"chat_id\":\"%d\"}", text, chatID))

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
