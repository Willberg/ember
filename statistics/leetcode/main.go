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

func helper(name string, start, end int) {
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
		for _, l := range rank.TotalRank {
			fmt.Printf("%s, score: %d, rank: %d, rankV2: %d\n", l.RealName, l.Score, l.Rank, l.RankV2)
		}
	}
}

func main() {
	helper("biweekly-contest-87", 1, 100)
}
