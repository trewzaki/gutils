package gutils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func LineOABoardcastMessages(channelAccessToken string, message []string) {
	messagePayload := map[string]interface{}{
		"messages": []map[string]interface{}{},
	}

	messages := []map[string]interface{}{}
	for _, v := range message {
		messages = append(messages, map[string]interface{}{
			"type": "text",
			"text": v,
		})
	}

	messagePayload["messages"] = messages
	jsonData, _ := json.Marshal(messagePayload)

	client := &http.Client{}
	req, _ := http.NewRequest("POST", "https://api.line.me/v2/bot/message/broadcast", bytes.NewBuffer(jsonData))

	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", channelAccessToken))

	res, _ := client.Do(req)
	ioutil.ReadAll(res.Body)
}
