# GeminiWhatsappBot

A WhatsApp bot with Google Gemini AI integration for intelligent message responses.

## Description

WhatsappBot is a Go application that creates a WhatsApp bot capable of automatically responding to messages using Google Gemini AI. The bot can handle multiple WhatsApp sessions, generate QR codes for authentication, and use AI to create contextual responses.

## Features

- 🤖 Automatic response to WhatsApp messages
- 🧠 Google Gemini AI integration for intelligent responses
- 📱 Support for multiple WhatsApp sessions
- 🔐 Secure credential management via `.env` file
- 📊 Detailed operation logging

## Prerequisites

- Go 1.16 or higher
- WhatsApp account
- Google Gemini API key (get it from [Google AI Studio](https://aistudio.google.com/))

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/GeminiWhatsappBot.git
   cd GeminiWhatsappBot
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Create a `.env` file in the project's root directory:
   ```
   GOOGLE_API_KEY=your_google_api_key_here
   ```

4. Create the necessary directories (if they don't exist):
   ```bash
   mkdir -p sessions assets
   ```

## Usage

### Start with default session:

```bash
go run . default
```

### Start with a custom session:

```bash
go run . mysession
```

On first launch, the bot will generate a QR code that you'll need to scan with WhatsApp to authenticate. The QR code is also saved as an image in the `assets` folder.

## Project Structure

- `main.go`: Application entry point and session management
- `utils.go`: Utility functions for AI response generation and message sending
- `events.go`: Event handlers for WhatsApp messages
- `.env`: Configuration file for API keys (not tracked by Git)
- `sessions/`: Directory for WhatsApp session databases
- `assets/`: Directory for generated files such as QR codes

## How It Works

1. The bot connects to WhatsApp using the whatsmeow library
2. When receiving a message, it processes the content using Google Gemini AI
3. It generates a contextually appropriate response
4. The response is sent back to the original sender

## Customization

You can customize the bot by modifying the `generateGeminiResponse` function in `utils.go` to create different prompts or change the language of responses. Currently, responses are generated with a 100-character limit.

## Security

- The `.env` file should be excluded from the Git repository to protect API credentials
- WhatsApp sessions are stored locally to ensure privacy
- Authentication tokens are stored in protected SQLite databases

## Dependencies

- [go-whatsmeow](https://github.com/tulir/whatsmeow) - Go library for WhatsApp
- [Google Gemini API](https://github.com/googleapis/go-genai) - AI intelligence API
- [go-qrcode](https://github.com/yeqown/go-qrcode) - QR code generation
- [godotenv](https://github.com/joho/godotenv) - Environment variable loading
- [go-sqlite3](https://github.com/mattn/go-sqlite3) - SQLite database driver

## License

This project is available under the MIT License.

## Contributing

Contributions are welcome! Feel free to submit issues or pull requests to improve this bot.
