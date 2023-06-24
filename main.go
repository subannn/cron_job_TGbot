package main

import (
	"time"
	
	//"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	handler "github.com/subannn/TelegramBot/handler"
	tgBot "github.com/subannn/TelegramBot/tgBot"
)

func runTicker(ann handler.Announcement) {
	ticker := time.NewTicker(ann.AnnouncementData)
    for {
        select {
        case <-ticker.C:
			sendToAll(ann)
        }
    }
}

func sendToAll(ann handler.Announcement) {
	msg := tgbotapi.NewForward(0, ann.ChatID, int(ann.MessageID))
	for id := range handler.Users_ID {
		msg.ChatID = id
		tgBot.Bot.Send(msg)
	}
}

func чоБля(chAnnouncement chan handler.Announcement) {
	for {
		select {
		case ann := <-chAnnouncement:
			runTicker(ann)
		}
	}
}

func chat_messages(updates tgbotapi.UpdatesChannel) {
	chAnnouncement := make(chan handler.Announcement)
	go чоБля(chAnnouncement)

	for update := range updates {
		if update.Message != nil { // If we got a message
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID
			go handler.Handle(&msg, &chAnnouncement)			
		}
		
	}
}
func main() {
	tgBot.StartBot()

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := tgBot.Bot.GetUpdatesChan(u)

	chat_messages(updates)
}