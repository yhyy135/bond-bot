package main

import (
	"log"
	"time"

	workingday "github.com/Admingyu/go-workingday"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	var (
		args Args
		conf Config
	)

	args.ReadFlags()
	conf.ReadConfig(args.Path)

	bot, err := tgbotapi.NewBotAPI(conf.Token)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Authorized account", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 10

	var timezone, _ = time.LoadLocation("Asia/Shanghai")

	var (
		prevDay int
		workDay bool
	)

	currentDay := time.Now().In(timezone).Day()
	if prevDay != currentDay {
		workDay, _ = workingday.IsWorkDay(
			time.Now().In(timezone), "CN",
		)

		prevDay = currentDay
	}

	if workDay {
		bonds := BondParser(BondData())
		MessageSender(
			bot, conf.ChatId, BondFilter(bonds),
		)
	}
}
