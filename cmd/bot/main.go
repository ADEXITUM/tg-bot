package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	telegram "tg-bot/pkg/telegram/actions"
	"tg-bot/pkg/telegram/types"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func setWebhook(bot *tgbotapi.BotAPI) {
	webHookInfo := tgbotapi.NewWebhook(fmt.Sprintf("https://29d4-159-255-6-200.eu.ngrok.io/5845449188:AAGPf7Vc14jCbDo5XbjN8eNSghJ4qObWP4c"))
	_, err := bot.SetWebhook(webHookInfo)
	if err != nil {
		fmt.Println("error while setting up webhook", err)
	}
}

func main() {
	NewBot := telegram.InitBot()
	setWebhook(NewBot)

	fmt.Println("setWebhook done")
	message := func(w http.ResponseWriter, r *http.Request) {
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

		switch botCommand {
		case "/start":
			msg := tgbotapi.NewMessage(int64(chatGroup), "Command /start")
			NewBot.Send(msg)
		case "/hello":
			msg := tgbotapi.NewMessage(int64(chatGroup), "Command /hello")
			NewBot.Send(msg)
		default:
			msg := tgbotapi.NewMessage(int64(chatGroup), "idk this command")
			NewBot.Send(msg)
		}

	}

	http.HandleFunc("/", message)

	fmt.Println("Starting http server")
	msg := tgbotapi.NewMessage(572377674, "Bot is about to start")
	NewBot.Send(msg)

	log.Fatal(http.ListenAndServe("localhost:8443", nil))
}

// func main() {
// 	bot, err := tgbotapi.NewBotAPI("5845449188:AAGPf7Vc14jCbDo5XbjN8eNSghJ4qObWP4c")
// 	if err != nil {
// 		log.Panic(err)
// 	}

// 	bot.Debug = true

// 	db, err := initDB()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	chatIDRepository := boltdb.NewRepository(db)

// 	telegramBot := telegram.NewBot(bot, chatIDRepository)
// 	if err := telegramBot.Start(); err != nil {
// 		log.Fatal(err)
// 	}

// }

// func initDB() (*bolt.DB, error) {
// 	db, err := bolt.Open("bot.db", 0600, nil)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if err := db.Update(func(tx *bolt.Tx) error {
// 		_, err := tx.CreateBucketIfNotExists([]byte(storage.Name))
// 		if err != nil {
// 			return err
// 		}
// 		_, err = tx.CreateBucketIfNotExists([]byte(storage.ChatID))
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	}); err != nil {
// 		return nil, err
// 	}

// 	return db, nil
// }
