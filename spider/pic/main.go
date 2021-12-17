package main

import (
	"container/list"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"
)

var (
	crawlFile = flag.String("f", "", "抓取图片链接的文件")
	savePath  = flag.String("p", "", "抓取图片存放路径")
	dirName   = flag.String("d", "", "抓取图片存放文件夹名")
)

func main() {
	flag.Parse()

	if "" == *crawlFile || "" == *savePath || "" == *dirName {
		fmt.Println("参数不对")
		return
	}

	if !isExists(*crawlFile) {
		fmt.Println("抓取图片链接的文件不存在")
		return
	}
	sp := path.Join(*savePath, *dirName)
	if sp == *crawlFile {
		fmt.Println("抓取图片链接的文件与存放的目录名冲突")
		return
	}
	if !isExists(sp) {
		err := os.MkdirAll(sp, 0755)
		if err != nil {
			fmt.Printf("创建目录%s失败", sp)
			return
		}
	}
	crawl(*crawlFile, sp)
}

func isExists(s string) bool {
	if _, err := os.Stat(s); os.IsNotExist(err) {
		return false
	}
	return true
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
