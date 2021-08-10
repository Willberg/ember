package main

import (
	"flag"
	"fmt"
	"github.com/tidwall/pretty"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
	url2 "net/url"
)

type env struct {
	Ak  string `yaml:"apiKey"`
	Ask string `yaml:"apiSecretKey"`
	Bt  string `yaml:"bearerToken"`
	At  string `yaml:"accessToken"`
	Ats string `yaml:"accessTokenSecret"`
}

var fp = flag.String("p", "/home/john/mine/workplace/go/ember/api/twitter/env/keys.yml", "配置文件路径")

func main() {
	flag.Parse()

	e := env{}
	data, err := ioutil.ReadFile(*fp)
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(data, &e)
	if err != nil {
		log.Fatal(err)
	}

	encodeStr := url2.QueryEscape("epidemic situation")
	url := "https://api.twitter.com/2/tweets/search/recent?query=" + encodeStr
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	bt := "Bearer " + e.Bt
	req.Header.Add("Authorization", bt)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s\n", pretty.Pretty(body))
}
