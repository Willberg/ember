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
	"sort"
	"strings"
	"sync"
	"time"
)

const (
	IMG = "bmp|jpg|png|tif|gif|pcx|tga|exif|fpx|svg|psd|cdr|pcd|dxf|ufo|eps|ai|raw|WMF|webp|avif|apng"
)

var (
	singleProject  = flag.Bool("s", true, "是否单独项目")
	fpath          = flag.String("p", "", "项目位置")
	rnum           = flag.Int("n", 10, "并发数")
	expDir         = flag.String("xd", "", "去除目录")
	expFile        = flag.String("xf", "", "去除文件")
	expComment     = flag.Bool("xc", true, "是否去除注释")
	checkFile      = flag.String("f", "", "待查文件类型")
	closeUnRegular = flag.Bool("c", true, "是否去除无文件类型的文件")
	checkFileSet   = ""
	sema           = make(chan struct{}, *rnum)
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
	if len(*fpath) == 0 {
		log.Fatal("请选择目录")
	}
	checkFileSet = strings.ToLower(strings.TrimSpace(*checkFile))

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

			// 去除图片
			if isFilter(IMG, f.Name()) {
				continue
			}

			// 根据条件去除无文件类型的文件
			if *closeUnRegular && !strings.Contains(f.Name()[:len(f.Name())-1], ".") {
				continue
			}
		}

		n.Add(1)
		project := f.Name()
		if *singleProject {
			project = getProjectName(*fpath)
		}
		dir := filepath.Join(*fpath, f.Name())
		switch mode := f.Mode(); {
		case mode.IsRegular():
			go wakFile(project, dir, &n, fileInfos)
		case mode.IsDir():
			go wakDir(project, dir, &n, fileInfos)
		case mode&os.ModeSymlink != 0:
			go func() {
				abPath, err := filepath.EvalSymlinks(dir)
				if err != nil {
					fmt.Printf("eval symlinks error %s err, err: %v\n", dir, err)
					return
				}
				if !strings.Contains(abPath, project) {
					n.Add(1)
					go wakDir(project, dir, &n, fileInfos)
				}
			}()
		default:
			go func() {
				fmt.Printf("unsovled link: %s \n", dir)
			}()
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
				// 单独的break，只能跳出select语句，无法跳出for 和 select双层语句
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
	// 按项目名排序
	proSlice := make([]string, 0, len(result))
	for k := range result {
		proSlice = append(proSlice, k)
	}
	sort.Strings(proSlice)

	for _, k1 := range proSlice {
		fmt.Printf("%s: \n", k1)
		var total int64
		for _, v2 := range result[k1] {
			total += v2
		}

		lanSlice := make([]string, 0, len(result[k1]))
		for k2 := range result[k1] {
			lanSlice = append(lanSlice, k2)
		}
		// 按行数排名, 然后按文件名排名
		sort.Slice(lanSlice, func(i, j int) bool {
			li, lj := result[k1][lanSlice[i]], result[k1][lanSlice[j]]
			if li != lj {
				return li > lj
			}
			return lanSlice[i] < lanSlice[j]
		})

		for _, k2 := range lanSlice {
			v2 := result[k1][k2]
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
		switch mode := entry.Mode(); {
		case mode.IsRegular():
			if isFilter(*expFile, entry.Name()) {
				continue
			}

			// 去除图片
			if isFilter(IMG, entry.Name()) {
				continue
			}

			// 如果待查文件类型存在，则只考虑对应类型的文件
			if isFilterFileWithSuffix(entry.Name()) {
				continue
			}

			// 根据条件去除无文件类型的文件
			if *closeUnRegular && !strings.Contains(entry.Name()[:len(entry.Name())-1], ".") {
				continue
			}

			readFile(project, subdir, fileInfos)
		case mode.IsDir():
			if isFilter(*expDir, entry.Name()) {
				continue
			}

			n.Add(1)
			go wakDir(project, subdir, n, fileInfos)
		case mode&os.ModeSymlink != 0:
			if isFilter(*expDir, entry.Name()) {
				continue
			}

			abPath, err := filepath.EvalSymlinks(subdir)
			if err != nil {
				fmt.Printf("eval symlinks error %s err, err: %v\n", subdir, err)
				continue
			}

			// 不在项目里说明之前并未统计，因此需要统计一次
			if !strings.Contains(abPath, project) {
				n.Add(1)
				go wakDir(project, dir, n, fileInfos)
			}
		default:
			fmt.Printf("unsovled link: %s \n", dir)
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
	isComment := false
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
		/*dfs*/
		/*
			dfs
		*/
		codeLine := string(line)
		codeLine = strings.TrimSpace(codeLine)
		if isPrefix != true && len(codeLine) > 0 {
			if *expComment {
				if !isComment {
					if strings.HasPrefix(codeLine, "//") {
						////fmt.Println(codeLine)
						continue
					}

					if strings.HasPrefix(codeLine, "/*") {
						if strings.HasSuffix(codeLine, "*/") {
							//fmt.Println(codeLine)
							continue
						} else {
							//fmt.Println(codeLine)
							isComment = true
						}
					} else {
						lineNum++
					}
				} else {
					//fmt.Println(codeLine)
					if strings.HasSuffix(codeLine, "*/") {
						isComment = false
						//fmt.Println(codeLine)
					}
				}
			} else {
				lineNum++
			}
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
