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
		"unread_count": "5",
	}
	valuesToAdd := map[string][]string{
		"Interests": []string{"Tech", "Sports"},
	}
	leanplum_users.Start(config, values)
	attributes := leanplum_users.NewAttributeContent("1")
	attributes.SetAttribute("UserAttributes", values)
	attributes.SetSliceOfAttributes("UserAttributesToAdd", valuesToAdd)
	slice, _ := leanplum_users.SendAttributes(config, attributes)

	for _, v := range slice {
		success, err := v.CheckErrors()

		fmt.Println(success, err)
	}
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
