package main

import (
	"ember/fs"
	. "ember/statistics/leetcode/contest/common"
	"fmt"
	"os"
	"time"
)

func main() {
	dir, _ := os.Getwd()
	dir += "/statistics/leetcode/contest/ide/"
	con, ok := fs.ReadJson(dir+"contest.json", &Contest{})
	if !ok {
		fmt.Errorf("%v\n", con)
		return
	}
	contest := con.(*Contest)
	path := fmt.Sprintf(dir+"%s.txt", contest.Name)
	startTime, _ := time.ParseInLocation("2006-01-02 15:04:05", contest.DateTime, time.Local)
	Process(contest.Name, path, contest.Start, contest.End, startTime)
}
