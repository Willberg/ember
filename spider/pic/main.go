package main

import (
	"container/list"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"
)

func main() {
	crawl("/home/john/Pictures/图库/测试", "/home/john/Pictures/图库")
}

func crawl(linkPath, saveDir string) {
	l := getUrls(linkPath)
	i := 1
	for p := l.Front(); p != nil; p = p.Next() {
		link := p.Value.(string)
		fmt.Printf("%s 抓取中\n", link)
		bs := crawlBinFile(link)
		name := fmt.Sprintf("%d", i)
		saveToFile(saveDir, name, bs)
		i++
	}
}

func getUrls(p string) *list.List {
	bytes, err := ioutil.ReadFile(p)
	if err != nil {
		os.Exit(1)
	}

	retList := list.New()
	for _, line := range strings.Split(string(bytes), "\n") {
		retList.PushBack(line)
	}

	return retList
}

func crawlBinFile(link string) *[]byte {
	resp, err := http.Get(link)
	if err != nil {
		fmt.Printf("%s error, reason: %s", link, err)
		return nil
	}

	if resp.StatusCode != 200 {
		fmt.Printf("%s error, reason: %s", link, resp.StatusCode)
		return nil
	}

	bs, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Printf("%s error, reason: %s", link, err)
		return nil
	}
	return &bs
}

func saveToFile(saveDir, name string, bs *[]byte) {
	p := path.Join(saveDir, name)
	err := ioutil.WriteFile(p, *bs, 0660)
	if err != nil {
		fmt.Printf("%s error, reason: %s", saveDir, err)
	}
}
