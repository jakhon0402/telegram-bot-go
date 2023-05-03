package main

import (
	"fmt"
	tele "gopkg.in/telebot.v3"
	"time"
)

func main() {
	pref := tele.Settings{
		Token:  "6134899031:AAHMLclPqnGJLY4pS3_PZDZqFvp6mEF5TJg",
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}
	b, err := tele.NewBot(pref)
	if err != nil {
		panic(err)
	}

	b.Handle("/start", func(c tele.Context) error {
		fmt.Println(c.Chat().ID)
		return c.Send("Assalomu alaykum!")
	})

	b.Start()
}
