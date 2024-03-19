package main

import (
	"github.com/orgs/KKGo-Software-engineering/fun-ex-git/contact"
	"github.com/orgs/KKGo-Software-engineering/fun-ex-git/notification"
	"github.com/orgs/KKGo-Software-engineering/fun-ex-git/toggle"
)

func main() {
	isEmailEnabled := false
	toggle := toggle.New(isEmailEnabled)

	contact := contact.New("han_f50@hotmail.com", "0812345678")
	notification := notification.New(toggle)
	notification.Send("Hello, Go!", contact)
}
