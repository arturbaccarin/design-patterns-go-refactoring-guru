package main

import "fmt"

/*
You're building a notification system that can send messages
through different channels like Email, SMS, and Push notifications.
You need to create a system that can generate these notifications based
on user input, but without directly instantiating the classes for each
type of notification. The goal is to use the Factory Method pattern
to allow the system to dynamically create the appropriate notification
object depending on the channel selected by the user.
*/

type Notification interface {
	Send()
}

func NewNotification(channel string) Notification {
	switch channel {
	case "email":
		return &EmailNotification{}
	case "sms":
		return &SMSNotification{}
	case "push":
		return &PushNotification{}
	default:
		return nil
	}
}

type EmailNotification struct {
}

func (EmailNotification) Send() {
	fmt.Println("Sending email notification")
}

type SMSNotification struct {
}

func (SMSNotification) Send() {
	fmt.Println("Sending SMS notification")
}

type PushNotification struct {
}

func (PushNotification) Send() {
	fmt.Println("Sending push notification")
}

func main() {
	notification := NewNotification("email")
	notification.Send()
}
