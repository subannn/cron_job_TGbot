package handler

import (
	"sync"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	tgBot "github.com/subannn/TelegramBot/tgBot"
)

func RunTicker(ann Announcement, Users_ID *map[int64]bool, Mut *sync.Mutex) {
	msg := tgbotapi.NewForward(0, ann.ChatID, int(ann.MessageID))
	ticker := time.NewTicker(ann.AnnouncementData)
    <-ticker.C
	Mut.Lock()
	defer Mut.Unlock()
	for id := range (*Users_ID) {
		msg.ChatID = id
		tgBot.Bot.Send(msg)
	}
	return
}

func Чо(chAnnouncement *chan Announcement, Users_ID *map[int64]bool, Mut *sync.Mutex) {
	for {
		select {
		case ann := <-*chAnnouncement:
			go RunTicker(ann, Users_ID, Mut)
		}
	}
}