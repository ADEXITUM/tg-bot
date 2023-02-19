package chatsonic

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"tg-bot/pkg/chatSonic/consts"
)

func CreateChatSonicMsg(text string) (response string) {
	url := consts.CS_URL

	payload := strings.NewReader(fmt.Sprintf("{\"enable_google_results\":true,\"enable_memory\":false,\"input_text\":\"%s\"}", text))

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		fmt.Println("Error while making new request ", err)
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("X-API-KEY", consts.API_KEY)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error while doing request: ", err)
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var responseText consts.Response

	err = json.Unmarshal(body, &responseText)
	if err != nil {
		fmt.Println("Error while unmarshaling msg")
	}

	return string(responseText.Message)
}
