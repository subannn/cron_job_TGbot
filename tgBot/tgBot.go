package tgBot

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var Bot *tgbotapi.BotAPI

func StartBot(){
	key := os.Getenv("BOT_KEY") 
	if(key == "") {
		log.Fatal("No key")
	}

 	b, err := tgbotapi.NewBotAPI(key)

	if err != nil {
		log.Panic(err)
	}
	
	b.Debug = true

	Bot = b

	log.Printf("Authorized on account %s", b.Self.UserName)
}
		