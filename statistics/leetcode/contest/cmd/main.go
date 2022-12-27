package main

import (
	. "ember/statistics/leetcode/contest/common"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

var (
	fp    = flag.String("p", "/Users/bill/Desktop/leetcode", "存储目录")
	name  = flag.String("n", "", "周赛名字")
	dt    = flag.String("d", "", "周赛开始时间")
	start = flag.Int("s", 1, "起始页")
	end   = flag.Int("e", 100, "终止页")
)

func main() {
	flag.Parse()
	if len(*fp) == 0 {
		log.Fatal("请填写存放目录")
	}
	if s, err := os.Stat(*fp); err != nil || !s.IsDir() {
		log.Fatal("目录不存在")
	}
	if len(*name) == 0 {
		log.Fatal("请填写周赛名字")
	}
	startTime, err := time.ParseInLocation("2006-01-02 15:04:05", *dt, time.Local)
	if err != nil {
		log.Fatal("请填写正确的周赛开始时间")
	}
	if (*fp)[len(*fp)-1:] == string(os.PathSeparator) {
		*fp = (*fp)[:len(*fp)-1]
	}
	path := fmt.Sprintf("%s/%s.txt", *fp, *name)
	Process(*name, path, *start, *end, startTime)
}
