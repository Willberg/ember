package main

import (
	"ember/api/twitter/common"
	"flag"
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"log"
)

var fp = flag.String("p", "/home/john/mine/workplace/go/ember/api/twitter/env/keys.yml", "配置文件路径")

func main() {
	flag.Parse()
	e := common.Parse(*fp)

	config := oauth1.NewConfig(e.Ak, e.Ask)
	token := oauth1.NewToken(e.At, e.Ats)
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	// Send a Tweet
	tweet, _, err := client.Statuses.Update("hello, world.", nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("text: %s\n", tweet.Text)
}
