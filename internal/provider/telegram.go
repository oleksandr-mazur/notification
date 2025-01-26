package provider

// https://github.com/go-telegram-bot-api/telegram-bot-api
// https://go-telegram-bot-api.dev/

import (
	"context"
	"log/slog"
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
		slog.Error(err.Error())
		panic(err)
	}

	bot.Debug = false

	slog.Info("Authorized on account " + bot.Self.UserName)
	message := tgbotapi.NewMessage(clientID, msg)

	bot.Send(message)
	slog.Info("Send message to", slog.Any("client", clientID))

	return nil
}
