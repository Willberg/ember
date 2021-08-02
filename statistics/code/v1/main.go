package main

import (
	"bufio"
	"container/list"
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

/**
  统计项目中每种语言的行数和占比,使用到协程同步WaitGroup,锁，并发,递归以及基本的map操作
*/
var (
	singleProject = flag.Bool("s", false, "是否单独项目")
	fpath         = flag.String("p", "/home/john/mine/workplace/go", "项目位置")
	rnum          = flag.Int("n", 10, "并发数")
	expDir        = flag.String("xd", "org|com|bin|img|vendor", "去除目录")
	expFile       = flag.String("xf", "problem|netcat|findlinks", "去除文件")
	tl            = new(fiList)
	me            = new(memo)
	n             sync.WaitGroup
)

type stat struct {
	lineNum int64
}

type memo struct {
	cache map[string]map[string]stat
	mu    sync.Mutex
}

func (m *memo) put(project string, language string, lineNum int64) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.cache == nil {
		m.cache = make(map[string]map[string]stat)
	}

	res, ok := m.cache[project]
	if !ok {
		m.cache[project] = make(map[string]stat)
		s := stat{
			lineNum: lineNum,
		}
		m.cache[project][language] = s
	} else {
		oldStat, ok := res[language]
		if !ok {
			s := stat{
				lineNum: lineNum,
			}
			m.cache[project][language] = s
		} else {
			oldStat.lineNum += lineNum
			m.cache[project][language] = oldStat
		}
	}
}

type listItem struct {
	project string
	link    string
}

type fiList struct {
	li *list.List
	mu sync.Mutex
}

func (l *fiList) getOne() *listItem {
	l.mu.Lock()
	defer l.mu.Unlock()
	if e := l.li.Front(); e != nil {
		l.li.Remove(e)
		return e.Value.(*listItem)
	}

	return nil
}

func (l *fiList) putOne(item *listItem) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.li == nil {
		l.li = list.New()
	}
	l.li.PushBack(item)
}

func (l *fiList) putAll(items *list.List) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.li == nil {
		l.li = list.New()
	}
	for e := items.Front(); e != nil; e = e.Next() {
		l.li.PushBack(e.Value)
	}
}

func main() {
	t := time.Now()
	flag.Parse()

	if *singleProject {
		if (*fpath)[len(*fpath)-1:] == string(os.PathSeparator) {
			*fpath = (*fpath)[:len(*fpath)-1]
		}
		pIdx := strings.LastIndex(*fpath, string(os.PathSeparator))
		pName := (*fpath)[pIdx+1:]
		l := getAllFiles(pName, *fpath)
		tl.putAll(l)
	} else {
		folds, err := ioutil.ReadDir(*fpath)
		if err != nil {
			log.Fatal("目录不存在")
		}

		for _, p := range folds {
			if p.Name()[:1] == "." {
				continue
			}

			fullName := *fpath + string(os.PathSeparator) + p.Name()
			if p.IsDir() {
				if isMatch, err := regexp.MatchString(*expDir, p.Name()); err == nil && isMatch {
					continue
				}
				l := getAllFiles(p.Name(), fullName)
				tl.putAll(l)
			} else {
				if isMatch, err := regexp.MatchString(*expFile, p.Name()); err == nil && isMatch {
					continue
				}
				item := new(listItem)
				item.project = p.Name()
				item.link = fullName
				tl.putOne(item)
			}
		}
	}

	//for e := tl.li.Front(); e != nil; e = e.Next() {
	//	v := e.Value.(*listItem)
	//	fmt.Printf("files %s, %s\n", v.project, v.link)
	//}

	// count project code lines
	i := 0
	for i < *rnum && i < tl.li.Len() {
		n.Add(1)
		go countFile()
		i++
	}
	n.Wait()

	// print records
	for k1, v1 := range me.cache {
		fmt.Printf("%s: \n", k1)
		var total int64
		for _, v2 := range v1 {
			total += v2.lineNum
			//fmt.Printf("%s: %d, %.2f \n", k2, v2.lineNum, v2.percent)
		}

		for k2, v2 := range v1 {
			percent := float64(v2.lineNum) / float64(total) * 100
			fmt.Printf("%s: %d, %.2f%% \n", k2, v2.lineNum, percent)
		}
		fmt.Println()
	}

	fmt.Printf("use time: %d ms", time.Since(t).Milliseconds())
}

// check file dir and list all file
func getAllFiles(project, dir string) *list.List {
	li := list.New()
	getAllFileFullName(project, dir, li)
	return li
}

func getAllFileFullName(project, dir string, l *list.List) {
	for _, entry := range dirents(dir) {
		if entry.Name()[:1] == "." {
			continue
		}

		subdir := filepath.Join(dir, entry.Name())
		if entry.IsDir() {
			if isMatch, err := regexp.MatchString(*expDir, entry.Name()); err == nil && isMatch {
				continue
			}
			getAllFileFullName(project, subdir, l)
		} else {
			if isMatch, err := regexp.MatchString(*expFile, entry.Name()); err == nil && isMatch {
				continue
			}
			item := new(listItem)
			item.project = project
			item.link = subdir
			l.PushBack(item)
		}
	}
}

func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "dirents err: %v", err)
		return nil
	}
	return entries
}

// count file code line and record
func countFile() {
	defer n.Done()
	for e := tl.getOne(); e != nil; e = tl.getOne() {
		//fmt.Printf("get file %s\n", e.link)

		pName := e.project
		link := e.link
		lanIdx := strings.LastIndex(link, ".")
		if lanIdx == -1 {
			lanIdx = strings.LastIndex(link, "/")
		}
		language := link[lanIdx+1:]

		f, err := os.Open(link)
		if err != nil {
			fmt.Fprintf(os.Stderr, "open file %s err, err: %v\n", link, err)
			f.Close()
			continue
		}

		var lineNum int64
		isRead := true
		buf := bufio.NewReader(f)
		for {
			line, isPrefix, err := buf.ReadLine()
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Fprintf(os.Stderr, "read file %s err, err: %v\n", link, err)
				isRead = false
				break
			}

			//fmt.Printf("line: %s\n", line)
			if isPrefix != true && len(line) > 0 {
				lineNum++
			}
		}
		f.Close()

		if isRead {
			me.put(pName, language, lineNum)
		}
	}
}
