package main

import (
    "errors"
    "fmt"
	"log"
    "net/mail"
    "strings"

    "github.com/pocketbase/pocketbase/tools/mailer"
)

func sendEmailIfContactInfoExists(data map[string]interface{}, recipientEmail string) error {
    _, emailExists := data["email"].(string)
    _, phoneExists := data["phone"].(string)

    if !emailExists && !phoneExists {
        return errors.New("No email or phone found in submission data")
    }

	if recipientEmail == "" {
		return errors.New("No forwarder email set")
	} 

    // Prepare the data list for the email body
    dataList := make([]string, 0, len(data))
    for key, value := range data {
        dataList = append(dataList, fmt.Sprintf("%s: %v", key, value))
    }
    dataContent := strings.Join(dataList, "\n")

    message := &mailer.Message{
        From: mail.Address{
            Address: "",
            Name:    "",
        },
        To:      mail.Address{[Address: recipientEmail]},
        Subject: "New submission received",
        HTML:    "<p>New submission information:</p><pre>" + dataContent + "</pre>",
    }

	return nil

    // return app.NewMailClient().Send(message)
}
