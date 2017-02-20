package main

import (
	"fmt"

	"github.com/agneum/leanplum-go-client"
	"github.com/agneum/leanplum-go-client/notifier"
	"github.com/agneum/leanplum-go-client/user"
)

func main() {
	// setUserAttributes()
	// sendMessage()
}

// Usage: setUserAttributes
func setUserAttributes() {
	config := leanplum.ReadConfig("config.toml")

	values := map[string]string{
		"action":         "setUserAttributes",
		"userId":         "1",
		"userAttributes": "{\"unread_count\":5}",
	}
	leanplum_users.Start(config, values)
}

// Usage: SendMessage
func sendMessage() {
	query := leanplum.ReadConfig("config.toml")
	values := map[string]string{
		"name": "John",
	}

	message := notifier.NewMessage("1", "4907543001825280")
	message.SetMessageValues(values)

	slice, _ := notifier.SendMessage(query, message)

	for _, v := range slice {
		success, err := v.CheckErrors()

		fmt.Println(success, err)
	}
}
