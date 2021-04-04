//Package main of ckselftask-telegram-bot a first contact telegram bot to checkin and checkout into ckselftask
package main

import (
	"fmt"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var mbot *tgbotapi.BotAPI

//init function
func init() {

}
func main() {
	var err error
	mbot, err = tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		panic(err)
	}

	mbot.Debug = true

	// Create a new UpdateConfig struct with an offset of 0. Offsets are used
	// to make sure Telegram knows we've handled previous values and we don't
	// need them repeated.
	updateConfig := tgbotapi.NewUpdate(0)

	// Tell Telegram we should wait up to 30 seconds on each request for an
	// update. This way we can get information just as quickly as making many
	// frequent requests without having to send nearly as many.
	updateConfig.Timeout = 30

	// Start polling Telegram for updates.
	updates := mbot.GetUpdatesChan(updateConfig)

	// Let's go through each update that we're getting from Telegram.
	for update := range updates {
		// Telegram can send many types of updates depending on what your Bot
		// is up to. We only want to look at messages for now, so we can
		// discard any other updates.
		if update.Message == nil {
			continue
		}

		//using our handler
		treatMessageFromBot(update.Message)
	}

}

func treatMessageFromBot(message *tgbotapi.Message) {

	if message.IsCommand() {
		treatBotCommand(message)
	} else {
		fmt.Println("Message from bot: " + message.Text)
	}
}

func treatBotCommand(message *tgbotapi.Message) {
	fmt.Println("Bot command received: " + message.Text)
	switch command := message.Text; command {
	case "/hello":
		fmt.Println("Return a hello message")
		responseMessageToBot(message, "Hello Mr. "+message.From.LastName)
	case "/weather":
		fmt.Println("Mr. is asking about the weather")
		treatWeatherCommand(message)
	default:
		// unknown message
		responseMessageToBot(message, "Sorry Mr. "+message.From.LastName+", I cannot do it")

	}
}

func responseMessageToBot(original *tgbotapi.Message, response string) {
	msg := tgbotapi.NewMessage(original.Chat.ID, response)

	msg.ReplyToMessageID = original.MessageID
	// Okay, we're sending our message off! We don't care about the message
	// we just sent, so we'll discard it.
	if _, err := mbot.Send(msg); err != nil {
		// Note that panics are a bad way to handle errors. Telegram can
		// have service outages or network errors, you should retry sending
		// messages or more gracefully handle failures.
		panic(err)
	}
}

func treatWeatherCommand(original *tgbotapi.Message) {
	responseMessageToBot(original, "it is sunny today")
}
