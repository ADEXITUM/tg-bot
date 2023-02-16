package main

import (
	"fmt"
	"log"
	"net/http"
	"tg-bot/pkg/telegram/actions"
)

func main() {
	actions.InitBot()
	actions.SetWebhook()

	fmt.Println("setWebhook done")

	http.HandleFunc("/", actions.HandleWebhook)

	fmt.Println("Starting http server")

	actions.SendMessage(572377674, "Bot is about to start")

	//TODO use certs
	log.Fatal(http.ListenAndServe("localhost:8443", nil))
}
