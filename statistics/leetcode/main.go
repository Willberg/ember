package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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
	Rank  int `json:"rank"`
	Score int `json:"score"`
	//FinishName    int    `json:"finish_time"`
	//GlobalRanking int    `json:"global_ranking"`
	//DataRegion    string `json:"data_region"`
	//AvatarUrl     string `json:"avatar_url"`
	RankV2 int `json:"rank_v2"`
}

func main() {
	for i := 1; i <= 49; i++ {
		res, err := http.Get(fmt.Sprintf("https://leetcode.cn/contest/api/ranking/biweekly-contest-86/?pagination=%d&region=local", i))
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
		ls := rank.TotalRank
		fmt.Println(len(ls))
		for _, l := range ls {
			fmt.Printf("%s, score: %d, rank: %d, rankV2: %d\n", l.RealName, l.Score, l.Rank, l.RankV2)
		}
	}
}
