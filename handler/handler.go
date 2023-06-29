package handler

import (
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	tgBot "github.com/subannn/TelegramBot/tgBot"
)

type Announcement struct {
	ChatID int64
	MessageID int64
	AnnouncementData time.Duration
}

func Handle(msg *tgbotapi.MessageConfig, chAnnouncement *chan Announcement, location *time.Location) {

	message := strings.Split(msg.Text, " ")
	if len(message) >= 4 && message[0] == "/new" && message[1] == "announcement" && IsTimeFormat(message[2] + " " + message[3]){
		var ann Announcement
		ann.ChatID = msg.ChatID

		timeInBishkek := time.Now().In(location)

		ann.MessageID = int64(msg.ReplyToMessageID)
		tm := StrToTime(message[2] + " " + message[3])

		ann.AnnouncementData = tm.Sub(timeInBishkek) - time.Hour * 6

		msg.Text = ann.AnnouncementData.String()
		
		if ann.AnnouncementData > time.Second * 5 {
			msg.Text = "You set announcement, time until: " + ann.AnnouncementData.String()
			tgBot.Bot.Send(msg)
			*chAnnouncement <- ann
		} else {
			msg.Text = "INCORRECT TIME"
			tgBot.Bot.Send(msg)
		}
		
	} else {
		msg.Text = "INCORRECT FORMAT"
		tgBot.Bot.Send(msg)
	}
	return
}