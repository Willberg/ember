package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

var (
	caPath = flag.String("p", "/Users/bill/shell/leetcode.cer", "请输入leetcode证书位置")
)

type Rank struct {
	realName string
	score    int
}

type RealName struct {
	Name string `json:"realName"`
}

type Contestant struct {
	Info RealName `json:"member"`
}

type Row struct {
	Score int        `json:"score"`
	User  Contestant `json:"contestant"`
}

type RankingBoard struct {
	Rows []Row `json:"rows"`
}

type Data struct {
	Board RankingBoard `json:"getRankingBoard"`
}

type ContestRank struct {
	RankData Data `json:"data"`
}

func main() {
	caCert, err := os.ReadFile(*caPath)
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			RootCAs: caCertPool,
		},
	}
	client := &http.Client{
		Transport: tr,
	}
	//ts := "{\"operationName\":\"rankingBoard\",\"variables\":{\"contestSlug\":\"2023-spring-solo\",\"from\":%d,\"size\":20},\"query\":\"query rankingBoard($contestSlug: String!, $from: Int!, $size: Int!) {\\n  getRankingBoard(contestSlug: $contestSlug, from: $from, size: $size) {\\n    total\\n    rows {\\n      ...rankingRow\\n      __typename\\n    }\\n    __typename\\n  }\\n}\\n\\nfragment rankingRow on RankingRow {\\n  score\\n  finishTime\\n  lastSubmissionId\\n  contestant {\\n    ... on SoloContestant {\\n      member {\\n        slug\\n        avatar\\n        realName\\n        __typename\\n      }\\n      __typename\\n    }\\n    ... on TeamContestant {\\n      captain {\\n        slug\\n        avatar\\n        realName\\n        __typename\\n      }\\n      members {\\n        slug\\n        avatar\\n        realName\\n        __typename\\n      }\\n      name\\n      slug\\n      slogan\\n      __typename\\n    }\\n    __typename\\n  }\\n  questionReport {\\n    questionId\\n    failedCount\\n    acSubmission {\\n      submissionId\\n      contestSubmissionId\\n      submittedAt\\n      submitterSlug\\n      lang\\n      __typename\\n    }\\n    __typename\\n  }\\n  __typename\\n}\\n\"}"
	ts := "{\"operationName\":\"rankingBoard\",\"variables\":{\"contestSlug\":\"2023-spring-solo\",\"from\":%d,\"size\":20},\"query\":\"query rankingBoard($contestSlug: String!, $from: Int!, $size: Int!) {\\n    getRankingBoard(contestSlug: $contestSlug, from: $from, size: $size) {\\n        rows {\\n            score\\n            contestant {\\n             ... on SoloContestant {\\n                 member {\\n                 realName\\n               }\\n            }\\n        }\\n    }\\n }\\n}\"}"
	ranks := make([]Rank, 0, 2500)
	for start := 0; start < 122; start++ {
		payload := strings.NewReader(fmt.Sprintf(ts, start*20))
		var contestRank ContestRank
		post(client, payload, &contestRank)
		for _, r := range contestRank.RankData.Board.Rows {
			rank := Rank{realName: r.User.Info.Name, score: r.Score}
			ranks = append(ranks, rank)
		}
		fmt.Printf("正在处理第%d组\n", start+1)
	}
	for _, r := range ranks {
		fmt.Printf("%s, %d\n", r.realName, r.score)
	}
}

func post(client *http.Client, payload *strings.Reader, rank interface{}) {
	url := "https://leetcode.cn/graphql/ranking"
	method := "POST"
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Accept-Language", "zh-CN,zh-Hans;q=0.9")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("Host", "leetcode.cn")
	req.Header.Add("Origin", "https://leetcode.cn")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.2 Safari/605.1.15")
	req.Header.Add("Connection", "keep-alive")

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(rank)
	if err != nil {
		fmt.Errorf("%v", err)
	}
}
