package provider

// https://github.com/go-telegram-bot-api/telegram-bot-api
// https://go-telegram-bot-api.dev/

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/go-telegram/bot"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SendTelegram(token string, clientID string, msg string) error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	b, err := bot.New(token)
	if err != nil {
		panic(err)
	}

	param := bot.SendMessageParams{ChatID: clientID, Text: msg}
	b.SendMessage(ctx, &param)
	return nil
}

func SendTelegram2(token string, clientID int64, msg string) error {

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)
	message := tgbotapi.NewMessage(clientID, "Test 2")

	bot.Send(message)

	return nil
}
