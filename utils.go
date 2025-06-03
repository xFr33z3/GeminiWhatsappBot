package main

import (
	"context"
	"fmt"

	"go.mau.fi/whatsmeow/proto/waE2E"
	"go.mau.fi/whatsmeow/types"
	"go.mau.fi/whatsmeow/types/events"
	"google.golang.org/genai"
	"google.golang.org/protobuf/proto"
)

var GoogleClient *genai.Client

var GoogleApiKey = ""

func generateGeminiResponse(prompt string) (string, error) {
	ctx := context.Background()
	model, _ := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  GoogleApiKey,
		Backend: genai.BackendGeminiAPI,
	})

	content := genai.Text(prompt)

	genConfig := &genai.GenerateContentConfig{
		MaxOutputTokens: 100,
	}

	resp, err := model.Models.GenerateContent(ctx, "gemini-1.5-flash", content, genConfig)
	if err != nil {
		return "", fmt.Errorf("error while generating response: %v", err)
	}

	if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
		return "", fmt.Errorf("no response was generated")
	}

	response := resp.Candidates[0].Content.Parts[0].Text
	return response, nil
}

func sendReply(to types.JID, replyToMsg *events.Message, text string) {
	messageToSend := text
	aiResponse, err := generateGeminiResponse("Write a friendly short response to this message: \"" + replyToMsg.Message.GetConversation() + "\". La risposta deve essere in italiano e non più lunga di 100 caratteri.")
	if err == nil {
		messageToSend = aiResponse
	} else {
		fmt.Printf("error using gemini, using default response: %v\n", err)
	}

	resp, err := client.SendMessage(context.Background(), to, &waE2E.Message{
		Conversation: proto.String(messageToSend),
	})
	if err != nil {
		fmt.Printf("Error while sending message: %v\n", err)
		return
	}
	fmt.Printf("Message sent with ID: %s\n", resp.ID)
}
