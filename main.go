package main

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/robfig/cron"
	"github/quin47/news/hackerNews"
	"log"
)

const botToken = ""
const myId = 0

func main() {

	c := cron.New()
	client, _ := tgbotapi.NewBotAPI(botToken)
	hackerNews := hackerNews.GetMarkDownFromHackNews()
	message := tgbotapi.NewMessage(myId, hackerNews)
	message.ParseMode = "markdown"
	c.AddFunc("0 0 0/8 * * *", func() {
		fmt.Println(message.Text)
		_, e := client.Send(message)
		if e != nil {
			log.Printf(" errors: %v", e)
		}

	})
	c.Start()
	select {}

}
