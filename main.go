package main

import (
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	handler "github.com/subannn/TelegramBot/handler"
	tgBot "github.com/subannn/TelegramBot/tgBot"
)

func chat_messages(updates tgbotapi.UpdatesChannel, chAnnouncement *chan handler.Announcement,  Users_ID *map[int64]bool, Mut *sync.Mutex, SuperUserID int, location *time.Location) {
	for update := range updates {
		if update.Message != nil { // If we got a message
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID
			if msg.ChatID == int64(SuperUserID) {
				go handler.Handle(&msg, chAnnouncement, location)
			} else {
				(*Users_ID)[msg.ChatID] = true
				msg.Text = "Wait for announcements"
				tgBot.Bot.Send(msg)
			}
		}
		
	}
}
func main() {
	tgBot.StartBot()
	SuperUserID, err := strconv.Atoi(os.Getenv("SUPER_USER_ID"))
	if err != nil {
		log.Fatal("Environment Variable is empty or not int: ", err)
	}
	location, err := time.LoadLocation("Asia/Bishkek")
	if err != nil {
		log.Fatal("Failed to load location:", err)
	}
	chAnnouncement := make(chan handler.Announcement)
	Users_ID := make(map[int64]bool)
	var Mut sync.Mutex
	
	Users_ID[int64(SuperUserID)] = true 

	go handler.Чо(&chAnnouncement, &Users_ID, &Mut)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := tgBot.Bot.GetUpdatesChan(u)

	chat_messages(updates, &chAnnouncement, &Users_ID, &Mut, SuperUserID, location)
}