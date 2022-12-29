package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"
)

var (
	t              = flag.Int("t", 2, "请输入类别 1-- 国内排名， 2-- 国际排名")
	caPath         = flag.String("p", "/Users/bill/shell/leetcode.cer", "请输入leetcode证书位置")
	myRank         = flag.Int("r", 34291, "请输入排名")
	myContestCount = flag.Int("c", 65, "请输入参赛数量")
)

type RealName struct {
	Name string `json:"realName"`
}

type RankNode struct {
	AttendedContestCount int      `json:"attendedContestCount"`
	CurrentRatingRanking int      `json:"currentRatingRanking"`
	User                 RealName `json:"user"`
}

type LocalRankingV2 struct {
	TotalUsers   int        `json:"totalUsers"`
	UserPerPage  int        `json:"userPerPage"`
	RankingNodes []RankNode `json:"rankingNodes"`
}

type LocalData struct {
	LocalRanking LocalRankingV2 `json:"localRankingV2"`
}

type LocalRank struct {
	Data LocalData `json:"data"`
}

type Profile struct {
	P RealName `json:"profile"`
}

type GlobalRankNode struct {
	Ranking              string  `json:"ranking"`
	CurrentGlobalRanking int     `json:"currentGlobalRanking"`
	User                 Profile `json:"user"`
}

type GlobalRanking struct {
	TotalUsers         int              `json:"totalUsers"`
	UserPerPage        int              `json:"userPerPage"`
	GlobalRankingNodes []GlobalRankNode `json:"rankingNodes"`
}

type GlobalRankData struct {
	GlobalR GlobalRanking `json:"globalRanking"`
}

type GlobalRank struct {
	Data GlobalRankData `json:"data"`
}

type Node struct {
	name         string
	rank         int
	contestCount int
}

func main() {
	start := time.Now().Unix()
	flag.Parse()
	if *myRank == -1 || *myContestCount == -1 {
		fmt.Println("请输入排名和参赛数量")
		return
	}
	var ts string
	if *t == 1 {
		ts = "{\"query\": \"{\\n  localRankingV2(page: %d) {\\n    page\\n    totalUsers\\n    userPerPage\\n    rankingNodes {\\n      attendedContestCount\\n      currentRatingRanking\\n      dataRegion\\n      isDeleted\\n      user {\\n        realName\\n      }\\n    }\\n  }\\n}\",\n  \"variables\": {}\n}"
	} else {
		ts = "{\"query\":\"{\\n  globalRanking(page: %d) {\\n    totalUsers\\n    userPerPage\\n    rankingNodes {\\n      currentGlobalRanking\\n  ranking\\n    dataRegion\\n      isDeleted\\n      user {\\n    profile {\\n    realName\\n      }\\n }\\n    }\\n  }\\n}\",\"variables\":{}}"
	}
	caCert, err := ioutil.ReadFile(*caPath)
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

	cnt := 0
	pageNo, q := 1, make([]Node, 0)
	payload := strings.NewReader(fmt.Sprintf(ts, pageNo))
	if *t == 1 {
		var rank LocalRank
		post(client, payload, &rank)
		for _, r := range rank.Data.LocalRanking.RankingNodes {
			if r.AttendedContestCount >= *myContestCount {
				cnt++
				if r.CurrentRatingRanking > *myRank {
					q = append(q, Node{r.User.Name, r.CurrentRatingRanking, r.AttendedContestCount})
				}
			}
		}
		fmt.Printf("正在统计国内排名，页码：%d\n", pageNo)
		total, pageNum := rank.Data.LocalRanking.TotalUsers, rank.Data.LocalRanking.UserPerPage
		pages := (total + pageNum - 1) / pageNum
		for i := 2; i <= pages; i++ {
			payload := strings.NewReader(fmt.Sprintf(ts, i))
			post(client, payload, &rank)
			for _, r := range rank.Data.LocalRanking.RankingNodes {
				if r.AttendedContestCount >= *myContestCount {
					cnt++
					if r.CurrentRatingRanking > *myRank {
						q = append(q, Node{r.User.Name, r.CurrentRatingRanking, r.AttendedContestCount})
					}
				}
			}
			fmt.Printf("正在统计国内排名，页码：%d, 进度：%.4f\n", i, float64(i)/float64(pages)*100)
		}
		fmt.Printf("国内参赛较多，排名较低的数量：%d, 总数：%d, 比例：%0.2f\n", len(q), cnt, float64(len(q))/float64(cnt))
	} else {
		var rank GlobalRank
		post(client, payload, &rank)
		for _, r := range rank.Data.GlobalR.GlobalRankingNodes {
			rc := 0
			for _, s := range strings.Split(r.Ranking[1:len(r.Ranking)-1], ",") {
				v, _ := strconv.Atoi(strings.TrimSpace(s))
				if v > 0 {
					rc++
				}
			}
			if rc >= *myContestCount {
				cnt++
				if r.CurrentGlobalRanking > *myRank {
					q = append(q, Node{r.User.P.Name, r.CurrentGlobalRanking, rc})
				}
			}
		}
		fmt.Printf("正在统计国际排名，页码：%d\n", pageNo)
		total, pageNum := rank.Data.GlobalR.TotalUsers, rank.Data.GlobalR.UserPerPage
		pages := (total + pageNum - 1) / pageNum
		for i := 2; i <= pages; i++ {
			payload := strings.NewReader(fmt.Sprintf(ts, i))
			post(client, payload, &rank)
			for _, r := range rank.Data.GlobalR.GlobalRankingNodes {
				rc := 0
				for _, s := range strings.Split(r.Ranking[1:len(r.Ranking)-1], ",") {
					v, _ := strconv.Atoi(strings.TrimSpace(s))
					if v > 0 {
						rc++
					}
				}
				if rc >= *myContestCount {
					cnt++
					if r.CurrentGlobalRanking > *myRank {
						q = append(q, Node{r.User.P.Name, r.CurrentGlobalRanking, rc})
					}
				}
			}
			fmt.Printf("正在统计国际排名，页码：%d, 进度：%.4f\n", i, float64(i)/float64(pages)*100)
		}
		fmt.Printf("国际参赛较多，排名较低的数量：%d, 总数：%d, 比例：%0.2f\n", len(q), cnt, float64(len(q))/float64(cnt))
	}
	sort.Slice(q, func(i, j int) bool {
		if q[i].rank != q[j].rank {
			return q[i].rank < q[j].rank
		}
		return q[i].contestCount > q[j].contestCount
	})
	for _, r := range q {
		fmt.Printf("%s, 排名：%d, 参赛数量：%d\n", r.name, r.rank, r.contestCount)
	}
	fmt.Printf("用时：%d\n", time.Now().Unix()-start)
}

func post(client *http.Client, payload *strings.Reader, rank interface{}) {
	url := "https://leetcode.cn/graphql"
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
