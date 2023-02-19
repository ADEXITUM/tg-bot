package actions

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	storage "tg-bot/pkg/storage/mysql"
	"tg-bot/pkg/telegram/types"
)

func HandleWebhook(w http.ResponseWriter, r *http.Request) {
	text, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Print("error setting up handler")
	}

	var botText types.BotMessage

	err = json.Unmarshal(text, &botText)
	if err != nil {
		fmt.Println("Error while unmarshaling msg")
	}

	fmt.Println(fmt.Sprintf("%s", text))

	username := botText.Message.From.Username
	chatUser := botText.Message.From.Id
	chatGroup := botText.Message.Chat.Id
	messageID := botText.Message.Message_ID
	botCommand := strings.Split(botText.Message.Text, "@")[0]
	commandText := strings.Split(botText.Message.Text, " ")

	fmt.Println(username, chatUser, chatGroup, messageID, botCommand, commandText)

	storage.InsertUserInfo(chatUser, username)

	switch botCommand {
	case "/start":
		SendMessage(int64(chatGroup), types.START_MSG)
	default:
		InvokeChatSonic(int64(chatGroup), botCommand)
	}

}
