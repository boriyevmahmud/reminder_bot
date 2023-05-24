package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reminder_bot/config"
	"reminder_bot/models"
	"time"
)

func main() {

	fmt.Println("time", time.Now())

	for {
		for _, text := range config.Texts {
			SendMessage("Say — " + text)
			time.Sleep(time.Minute * 30)
		}
	}

}

var sendMessage = "sendMessage"

func SendMessage(txt string) {

	message := models.Message{
		ChatID:    config.CHANNEL_CHAT_ID,
		Text:      txt,
		ParseMode: config.ParseModeHTML,
	}

	jsonStr, err := json.Marshal(message)
	if err != nil {
		fmt.Println("ERROR while marshalling:=>> ", err, message)
		return
	}

	req, err := http.NewRequest("POST",
		fmt.Sprintf("%s%s/%s", config.TelegramBaseURL, config.BOT_TOKEN, sendMessage),
		bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Println("ERROR while declaring req: ", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error while making request to telegram client.Do", err)
		return
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error while making request to telegram ioutil.ReadAll", fmt.Sprintf("resp %s", string(body)))
		return
	}

	if res.StatusCode != 403 && res.StatusCode != 200 {
		fmt.Println("Error while making request to telegram", fmt.Sprintf("resp %s", string(body)))
		println("Token:", config.BOT_TOKEN)
	}

	defer res.Body.Close()
}
