package main

import (
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	handler "github.com/subannn/TelegramBot/handler"
	tgBot "github.com/subannn/TelegramBot/tgBot"
)



func runTicker(ann *handler.Announcement, Users_ID *handler.SafeSet) {
	Users_ID.Mut.Lock()
	defer Users_ID.Mut.Unlock()
	msg := tgbotapi.NewForward(0, ann.ChatID, int(ann.MessageID))
	ticker := time.NewTicker(ann.AnnouncementData)
    for {
        select {
		case <-ticker.C:
			msg.ChatID = ann.ChatID
			tgBot.Bot.Send(msg)
			return
        }
    }
}

func чо(chAnnouncement *chan handler.Announcement, Users_ID *handler.SafeSet) {
	for {
		select {
		case ann := <-*chAnnouncement:
			runTicker(&ann, Users_ID)
		}
	}
}


func chat_messages(updates tgbotapi.UpdatesChannel) {
	chAnnouncement := make(chan handler.Announcement)
	var Users_ID handler.SafeSet
	go чо(&chAnnouncement, &Users_ID)

	for update := range updates {
		if update.Message != nil { // If we got a message
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID
			handler.Handle(&msg, &chAnnouncement, &Users_ID)			
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