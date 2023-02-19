package actions

import (
	"fmt"
	chatsonic "tg-bot/pkg/chatSonic/actions"
)

func InvokeChatSonic(chatID int64, text string) {
	fmt.Println("this is chat id in invokechatsonic: ", chatID)
	fmt.Println("this is text in invokechatsonic: ", text)
	response := chatsonic.CreateChatSonicMsg(text)
	fmt.Println("this is chat sonic response in invokechatsonic: ", response)
	SendMessage(chatID, response)
}
