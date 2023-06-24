package handler

import (
	"log"
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
func isTimeFormat(str string) bool{
    _, err := time.Parse("2006-01-02 15:04", str)
	if err != nil {
		return false
	}
	return true
}
func strToTime(str string) time.Time {
    t, err := time.Parse("2006-01-02 15:04", str)
	if err != nil {
		log.Panic(err)
	}
	return t
}

var Users_ID map[int64]bool
func Handle(msg *tgbotapi.MessageConfig, chAnnouncement *chan Announcement) {
	message := strings.Split(msg.Text, " ")
	if len(message) >= 4 && message[0] == "/new" && message[1] == "announcement" && isTimeFormat(message[2] + " " + message[3]){
		var ann Announcement
		ann.ChatID = msg.ChatID
		Users_ID[ann.ChatID] = true
		ann.MessageID = int64(msg.ReplyToMessageID)
		tm := strToTime(message[2] + " " + message[3])
		ann.AnnouncementData = time.Minute * 2
		isBefore := time.Now().Before(tm)
		if isBefore {
			*chAnnouncement <- ann
		}else {
			msg.Text = "INCORRECT TIME"
			tgBot.Bot.Send(msg)
		}
	} else {
		msg.Text = "INCORRECT FORMAT"
		tgBot.Bot.Send(msg)
	}
}