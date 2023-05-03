package main

import (
	"fmt"
	tele "gopkg.in/telebot.v3"
	"time"
)

type State string

const (
	START      State = "start"
	FIRSTNAME  State = "first_name"
	LASTNAME   State = "last_name"
	REGISTERED       = "registered"
)

type User struct {
	firstName string
	lastName  string
	state     State
}

//func (u User) FirstName() string {
//	return u.firstName
//}

func (u *User) SetFirstName(name string) {
	u.firstName = name
}

func (u *User) SetLastName(name string) {
	u.lastName = name
}

func (u *User) SetState(name State) {
	u.state = name
}

func main() {
	users := make(map[int64]*User)
	pref := tele.Settings{
		Token:  "6134899031:AAHMLclPqnGJLY4pS3_PZDZqFvp6mEF5TJg",
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}
	b, err := tele.NewBot(pref)
	if err != nil {
		panic(err)
	}

	b.Handle(tele.OnText, func(c tele.Context) error {
		currUser, ok := users[c.Chat().ID]
		if !ok {
			currUser = &User{state: START}
			users[c.Chat().ID] = currUser
		}
		fmt.Println(currUser)

		if c.Text() == "/start" {
			return c.Send("Assalomu alaykum! Ismingizni kiriting!")
		} else {
			if currUser.state == START {
				currUser.SetFirstName(c.Text())
				currUser.SetState(LASTNAME)
				return c.Send("Familiyangizni kiriting!")
			}
			if currUser.state == LASTNAME {
				currUser.SetLastName(c.Text())
				currUser.SetState(REGISTERED)
				return c.Send("Registered!")
			}
			if currUser.state == REGISTERED {
				return c.Send("Xush kelibsiz" + currUser.lastName + " " + currUser.firstName)
			}
			return c.Send("Invalid command")
		}
		//return c.Send(c.Text())
	})
	//b.Handle("/start", func(c tele.Context) error {
	//	fmt.Println(c.Chat().ID)
	//	return c.Send("Assalomu alaykum!")
	//})

	b.Start()
}
