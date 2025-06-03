package main

import (
	"fmt"

	"go.mau.fi/whatsmeow/types/events"
)

func eventHandler(evt interface{}) {
	switch v := evt.(type) {
	case *events.Message:
		fmt.Println("Received a message!", v.Message.GetConversation())

		if !v.Info.IsFromMe {
			fmt.Printf("Received message from %s: %s\n", v.Info.Sender.String(), v.Message.GetConversation())

			replyText, err := generateGeminiResponse(v.Message.GetConversation())
			if err != nil {
				sendReply(v.Info.Sender, v, "Error while generating response.")
			} else {
				sendReply(v.Info.Sender, v, replyText)
			}

		} else {
			fmt.Println("Message sent by myself. Let's ignore that.")
		}
	}
}
