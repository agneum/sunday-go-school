package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/agneum/leanplum-go-client"
)

func main() {

	query := leanplum.ReadConfig("config.toml")
	values := map[string]string{
		"expert_first_name":     "John",
		"expert_last_name":      "Doe",
		"expert_email":          "john@test.com",
		"client_first_name":     "Mike",
		"client_profil_pic_url": "",
		"message_content":       "Hi John!",
		"teep_id":               "123",
		"demand_id":             "111",
	}

	body, err := json.Marshal(leanplum.Message{Data: leanplum.MessageContent{Time: time.Now().Unix(), Values: values}})
	if err != nil {
		log.Fatalf("%v", err)
		return
	}

	params := map[string]string{
		"action":    "sendMessage",
		"userId":    "1",
		"messageId": "5891979810963456",
	}

	leanplum.Post(query, body, params)
}
