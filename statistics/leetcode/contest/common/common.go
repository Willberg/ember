package common

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

type Rank struct {
	//IsPast      bool   `json:"is_past"`
	//Submissions string `json:"submissions"`
	//Questions   string `json:"questions"`
	TotalRank []Leetcode `json:"total_rank"`
	//UserNum     int    `json:"user_num"`
}

type Leetcode struct {
	//ContestId     int    `json:"contest_id"`
	//UserName      string `json:"username"`
	//UserSlug      string `json:"user_slug"`
	RealName string `json:"real_name"`
	//CountryCode   string `json:"country_code"`
	//CountryName   string `json:"country_name"`
	Rank       int `json:"rank"`
	Score      int `json:"score"`
	FinishTime int `json:"finish_time"`
	//GlobalRanking int    `json:"global_ranking"`
	//DataRegion    string `json:"data_region"`
	//AvatarUrl     string `json:"avatar_url"`
	RankV2 int `json:"rank_v2"`
}

func Process(name, path string, start, end int, startTime time.Time) {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Errorf("%v", err)
		return
	}
	defer func() {
		err := file.Close()
		fmt.Errorf("%v", err)
		return
	}()
	for i := start; i <= end; i++ {
		res, err := http.Get(fmt.Sprintf("https://leetcode.cn/contest/api/ranking/%s/?pagination=%d&region=local", name, i))
		if err != nil {
			fmt.Errorf("%v", err)
			return
		}
		var rank Rank
		err = json.NewDecoder(res.Body).Decode(&rank)
		res.Body.Close()
		if err != nil {
			fmt.Errorf("%v", err)
		}
		sb := &strings.Builder{}
		for _, l := range rank.TotalRank {
			subTime := int(time.Unix(int64(l.FinishTime), 0).Sub(startTime).Seconds())
			useTime := fmt.Sprintf("%d时%d分%d秒", subTime/3600, subTime%3600/60, subTime%60)
			s := fmt.Sprintf("%s, score: %d, rank: %d, rankV2: %d, finishTime:%s\n", l.RealName, l.Score, l.Rank, l.RankV2, useTime)
			fmt.Print(s)
			sb.WriteString(s)
		}
		write := bufio.NewWriter(file)
		_, err = write.WriteString(sb.String())
		if err != nil {
			fmt.Errorf("%v", err)
		}
		err = write.Flush()
		if err != nil {
			fmt.Errorf("%v", err)
		}
	}
}

type Contest struct {
	Name     string `json:"name"`
	Start    int    `json:"start"`
	End      int    `json:"end"`
	DateTime string `json:"datetime"`
}
