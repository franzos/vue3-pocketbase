package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/mail"
	"strings"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/tools/mailer"
)

type EmailApiData struct {
	app            pocketbase.PocketBase
	data           map[string]interface{}
	recipientEmail string
}

type EmailApiResult struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	HTML    string `json:"html"`
}

func SendEmailIfContactInfoExists(api EmailApiData) (*EmailApiResult, error) {
	if api.recipientEmail == "" {
		return nil, errors.New("No recipient email set")
	}

	// Prepare the data list for the email body
	dataList := make([]string, 0, len(api.data))
	for key, value := range api.data {
		dataList = append(dataList, fmt.Sprintf("%s: %v", key, value))
	}
	dataContent := strings.Join(dataList, "\n")

	message := &mailer.Message{
		From: mail.Address{
			Address: api.app.Settings().Meta.SenderAddress,
			Name:    api.app.Settings().Meta.SenderName,
		},
		To:      []mail.Address{{Address: api.recipientEmail, Name: ""}},
		Subject: "New submission received",
		HTML:    "<p>New submission information:</p><pre>" + dataContent + "</pre>",
		Text:    "New submission information:\n" + dataContent,
	}

	err := api.app.NewMailClient().Send(message)
	if err != nil {
		log.Printf("Error sending email: %v", err)
		return nil, err
	} else {
		result := &EmailApiResult{
			From:    message.From.Address,
			To:      message.To[0].Address,
			Subject: message.Subject,
		}
		return result, nil
	}
}

func (er EmailApiResult) toJSON() string {
	res, err := json.Marshal(er)
	if err != nil {
		log.Printf("Error marshalling email result: %v", err)
		return ""
	}
	return string(res)
}
