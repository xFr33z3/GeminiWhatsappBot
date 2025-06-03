package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/yeqown/go-qrcode/v2"
	"github.com/yeqown/go-qrcode/writer/standard"

	_ "github.com/mattn/go-sqlite3"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store/sqlstore"
	waLog "go.mau.fi/whatsmeow/util/log"

	"google.golang.org/genai"
)

var client *whatsmeow.Client

func main() {
	args := os.Args
	if len(args) <= 1 {
		fmt.Println("Invalid arguments")
		os.Exit(0)
	}
	sessionName := args[1]

	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Cannot load GOOGLE_API_KEY environment variable - %v\n", err)
		os.Exit(1)
	}

	if GoogleApiKey == "" {
		GoogleApiKey = os.Getenv("GOOGLE_API_KEY")
		if GoogleApiKey == "" {
			fmt.Println("WARNING: GOOGLE_API_KEY empty. Fill that in the .env file.")
		} else {
			fmt.Println("GOOGLE_API_KEY loaded.")
		}
	}

	dbLog := waLog.Stdout("Database", "DEBUG", true)
	ctx := context.Background()
	container, err := sqlstore.New(ctx, "sqlite3", fmt.Sprintf("file:sessions/%v.db?_foreign_keys=on", sessionName), dbLog)
	if err != nil {
		panic(err)
	}
	deviceStore, err := container.GetFirstDevice(ctx)
	if err != nil {
		panic(err)
	}
	clientLog := waLog.Stdout("Client", "DEBUG", true)
	client = whatsmeow.NewClient(deviceStore, clientLog)
	client.AddEventHandler(eventHandler)

	var err2 error
	GoogleClient, err2 = genai.NewClient(ctx, &genai.ClientConfig{
		APIKey: GoogleApiKey,
	})
	if err2 != nil {
		fmt.Printf("Error while initializing Google AI: %v\n", err2)
	} else {
		fmt.Println("Client Google AI loaded successfully")
	}

	if client.Store.ID == nil {
		// No ID stored, new login
		qrChan, _ := client.GetQRChannel(context.Background())
		err = client.Connect()
		if err != nil {
			panic(err)
		}
		for evt := range qrChan {
			if evt.Event == "code" {
				fmt.Println("QR code:", evt.Code)

				qrc, err := qrcode.New(evt.Code)
				if err != nil {
					fmt.Printf("could not generate QRCode: %v", err)
					return
				}

				w, err := standard.New("./assets/repo-qrcode.jpeg")
				if err != nil {
					fmt.Printf("standard.New failed: %v", err)
					return
				}

				if err = qrc.Save(w); err != nil {
					fmt.Printf("could not save image: %v", err)
				}

			} else {
				fmt.Println("Login event:", evt.Event)
			}
		}
	} else {
		err = client.Connect()
		if err != nil {
			panic(err)
		}

	}

	// Listen to Ctrl+C
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	client.Disconnect()
}
