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
	singleProject = flag.Bool("s", true, "是否单独项目")
	fpath         = flag.String("p", "/home/john/mine/workplace/c", "项目位置")
	rnum          = flag.Int("n", 10, "并发数")
	expDir        = flag.String("xd", "tests|fail2ban|venv|bin|vendor|kernel_liteos_a", "去除目录")
	expFile       = flag.String("xf", "LICENSE|Dockerfile", "去除文件")
	checkFile     = flag.String("f", "h,c,h,hpp,hxx,cpp,cc,cxx,c++", "待查文件类型")
	checkFileSet  = strings.ToLower(*checkFile)
	sema          = make(chan struct{}, *rnum)
)

type item struct {
	project  string
	language string
	lineNum  int64
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
		// 如果是隐藏文件或文件夹，直接过滤
		if strings.HasPrefix(f.Name(), ".") {
			continue
		}

		// 如果待查文件类型存在，则只考虑对应类型的文件
		if !f.IsDir() && isFilterFileWithSuffix(f.Name()) {
			continue
		}

		if f.IsDir() {
			if isFilter(*expDir, f.Name()) {
				continue
			}
		} else {
			if isFilter(*expFile, f.Name()) {
				continue
			}
		}

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

func getFileSuffix(f string) string {
	idx := strings.LastIndex(f, ".")
	if idx > 0 {
		return f[idx+1:]
	} else {
		return ""
	}
}

func isFilterFileWithSuffix(f string) bool {
	if len(checkFileSet) > 0 {
		if suffix := getFileSuffix(f); len(suffix) > 0 {
			suffix = strings.ToLower(suffix)
			return !strings.Contains(checkFileSet, suffix)
		}
		return true
	}

	return false
}

func isFilter(t, f string) bool {
	if len(t) > 0 {
		if isMatch, err := regexp.MatchString(t, f); err == nil && isMatch {
			return true
		}
	}
	return false
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
		// 如果是隐藏文件或文件夹，直接过滤
		if strings.HasPrefix(entry.Name(), ".") {
			continue
		}

		subdir := filepath.Join(dir, entry.Name())
		if entry.IsDir() {
			if isFilter(*expDir, entry.Name()) {
				continue
			}

			n.Add(1)
			go wakDir(project, subdir, n, fileInfos)
		} else {
			if isFilter(*expFile, entry.Name()) {
				continue
			}

			// 如果待查文件类型存在，则只考虑对应类型的文件
			if isFilterFileWithSuffix(entry.Name()) {
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
