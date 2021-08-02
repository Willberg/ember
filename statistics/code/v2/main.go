package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"time"
)

var (
	singleProject = flag.Bool("s", false, "是否单独项目")
	fpath         = flag.String("p", "/home/john/mine/workplace/go", "项目位置")
	rnum          = flag.Int("n", 10, "并发数")
	expDir        = flag.String("xd", "org|com|bin|img|vendor", "去除目录")
	expFile       = flag.String("xf", "problem|netcat|findlinks", "去除文件")
	sema          = make(chan struct{}, *rnum)
)

type item struct {
	project  string
	language string
	lineNum  int64
}

type linkItem struct {
	project string
	link    string
}

// 通过通道来进行通信, 无锁并发；WaitGroup来同步;递归和map操作；效率高于v1
func main() {
	t := time.Now()
	flag.Parse()

	// 遍历目录
	var n sync.WaitGroup
	fileInfos := make(chan item)
	if (*fpath)[len(*fpath)-1:] == string(os.PathSeparator) {
		*fpath = (*fpath)[:len(*fpath)-1]
	}

	folds, err := ioutil.ReadDir(*fpath)
	if err != nil {
		log.Fatal("目录不存在")
	}
	for _, f := range folds {
		n.Add(1)
		project := f.Name()
		if *singleProject {
			project = getProjectName(*fpath)
		}
		dir := filepath.Join(*fpath, f.Name())
		if f.IsDir() {
			go wakDir(project, dir, &n, fileInfos)
		} else {
			go wakFile(project, dir, &n, fileInfos)
		}
	}

	// 关闭文件统计通道，结束下面的循环
	go func() {
		n.Wait()
		close(fileInfos)
	}()

	// 统计行数，并打印
	result := make(map[string]map[string]int64)
loop:
	for {
		select {
		case im, ok := <-fileInfos:
			if !ok {
				break loop
			}
			putResult(result, im)
		}
	}

	printResult(result)
	fmt.Printf("use time: %d ms", time.Since(t).Milliseconds())
}

func getProjectName(n string) string {
	pIdx := strings.LastIndex(n, string(os.PathSeparator))
	return n[pIdx+1:]
}

func putResult(result map[string]map[string]int64, im item) {
	res, ok := result[im.project]
	if !ok {
		result[im.project] = make(map[string]int64)
		res = result[im.project]
	}

	lineNum, ok := res[im.language]
	if !ok {
		res[im.language] = im.lineNum
	} else {
		lineNum += im.lineNum
		res[im.language] = lineNum
	}
}

func printResult(result map[string]map[string]int64) {
	// print records
	for k1, v1 := range result {
		fmt.Printf("%s: \n", k1)
		var total int64
		for _, v2 := range v1 {
			total += v2
			//fmt.Printf("%s: %d, %.2f \n", k2, v2.lineNum, v2.percent)
		}

		for k2, v2 := range v1 {
			percent := float64(v2) / float64(total) * 100
			fmt.Printf("%s: %d, %.2f%% \n", k2, v2, percent)
		}
		fmt.Println()
	}
}

func wakDir(project, dir string, n *sync.WaitGroup, fileInfos chan item) {
	defer n.Done()

	for _, entry := range dirents(dir) {
		if entry.Name()[:1] == "." {
			continue
		}

		subdir := filepath.Join(dir, entry.Name())
		if entry.IsDir() {
			if isMatch, err := regexp.MatchString(*expDir, entry.Name()); err == nil && isMatch {
				continue
			}

			n.Add(1)
			go wakDir(project, subdir, n, fileInfos)
		} else {
			if isMatch, err := regexp.MatchString(*expFile, entry.Name()); err == nil && isMatch {
				continue
			}
			readFile(project, subdir, fileInfos)
		}
	}
}

func wakFile(project, fp string, n *sync.WaitGroup, fileInfos chan item) {
	// 获取令牌
	sema <- struct{}{}
	// 释放令牌
	defer func() {
		<-sema
	}()

	defer n.Done()

	idx := strings.LastIndex(fp, string(os.PathSeparator))
	fname := fp[idx+1:]
	if fname[:1] == "." {
		return
	}

	readFile(project, fp, fileInfos)
}

func dirents(dir string) []os.FileInfo {
	// 获取令牌
	sema <- struct{}{}
	// 释放令牌
	defer func() {
		<-sema
	}()

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Printf("dirents err: %v", err)
		return nil
	}
	return entries
}

func readFile(project, link string, fileInfos chan item) {
	f, err := os.Open(link)
	if err != nil {
		fmt.Printf("open file %s err, err: %v\n", link, err)
		return
	}
	defer f.Close()

	var lineNum int64
	isRead := true
	buf := bufio.NewReader(f)
	for {
		line, isPrefix, err := buf.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("read file %s err, err: %v\n", link, err)
			isRead = false
			break
		}

		//fmt.Printf("line: %s\n", line)
		if isPrefix != true && len(line) > 0 {
			lineNum++
		}
	}

	if isRead {
		lanIdx := strings.LastIndex(link, ".")
		if lanIdx == -1 {
			lanIdx = strings.LastIndex(link, "/")
		}
		language := link[lanIdx+1:]
		im := item{
			project:  project,
			language: language,
			lineNum:  lineNum,
		}
		fileInfos <- im
	}
}
