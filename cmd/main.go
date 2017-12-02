package main

import (
  "log"
  "os"

  "github.com/nlopes/slack"
  "slackbot/config"
  "github.com/BurntSushi/toml"
)

func run(api *slack.Client) int {
    // websocket接続
    rtm := api.NewRTM()
    go rtm.ManageConnection()

    for {
        select {
        case msg := <-rtm.IncomingEvents:
            switch ev := msg.Data.(type) {
            case *slack.HelloEvent:
                log.Print("Hello Event")

            case *slack.MessageEvent:
                log.Printf("Message: %v\n", ev)
                rtm.SendMessage(rtm.NewOutgoingMessage("Hello world", ev.Channel))

            case *slack.InvalidAuthEvent:
                log.Print("Invalid credentials")
                return 1

            }
        }
    }
}

func main() {
  // config定義
  var cfg config.Config
  if _, err := toml.DecodeFile("development.toml", &cfg); err != nil {
    log.Println(err)
  }
  // slack.Client
  api := slack.New(cfg.SlackBot.Token)
  // run呼び出し、プロセス終了
  os.Exit(run(api))
}
