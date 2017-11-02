package main

import (
	"fmt"
  "os"
	"strings"

	"github.com/nlopes/slack"
)

func main() {

  token := os.Getenv("SLACK_TOKEN")
	api := slack.New(token)

	rtm := api.NewRTM()
	go rtm.ManageConnection()

Loop:
	for {
		select {
		case msg := <-rtm.IncomingEvents:
			fmt.Print("Event Received: ")
			switch ev := msg.Data.(type) {
			case *slack.ConnectedEvent:
				fmt.Println("Connection counter:", ev.ConnectionCount)

			case *slack.MessageEvent:
				fmt.Printf("Message: %v\n", ev)
				info := rtm.GetInfo()
				prefix := fmt.Sprintf("<@%s> ", info.User.ID)

				if ev.User != info.User.ID && strings.HasPrefix(ev.Text, prefix) {
					respond(rtm, ev, prefix)
				}

			case *slack.RTMError:
				fmt.Printf("Error: %s\n", ev.Error())

			case *slack.InvalidAuthEvent:
				fmt.Printf("Invalid credentials")
				break Loop

			default:
				//Take no action
			}
		}
	}
}

func respond(rtm *slack.RTM, msg *slack.MessageEvent, prefix string) {
	var response string
	text := msg.Text
	text = strings.TrimPrefix(text, prefix)
	text = strings.TrimSpace(text)
	text = strings.ToLower(text)

	acceptedGreetings := map[string]bool{
		"what's up": true,
		"hey":       true,
		"yo":         true,
	}
	acceptedHowAreYou := map[string]bool{
		"how's it going?": true,
		"how are ya?":     true,
		"feeling okay?":   true,
	}
  acceptedFav := map[string]bool{
		"favorite superhero?": true,
	}
  acceptedPlans := map[string]bool{
		"lets make plans": true,
	}


	if acceptedGreetings[text] {
		response = "BARLARP"
		rtm.SendMessage(rtm.NewOutgoingMessage(response, msg.Channel))
	} else if acceptedHowAreYou[text] {
		response = "BALL HUNGRY"
		rtm.SendMessage(rtm.NewOutgoingMessage(response, msg.Channel))
	} else if acceptedFav[text] {
		response = "BATMAN"
		rtm.SendMessage(rtm.NewOutgoingMessage(response, msg.Channel))
  } else if acceptedPlans[text] {
		response = "ILL ALWAYS RESPOND TO THIS UNLIKE SOME PEOPLE"
		rtm.SendMessage(rtm.NewOutgoingMessage(response, msg.Channel))
	}
}
