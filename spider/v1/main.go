package main

import (
	"bytes"
	"flag"
	"fmt"
	xmlpath "gopkg.in/xmlpath.v2"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	link = flag.String("l", "http://www.usgovernmentspending.com/", "抓取链接")
	xp   = flag.String("x", "//*[@id=\"broad_col3\"]/div[1]/table[1]/tbody/tr[1]/td[2]/b", "抓取规则")
	fp   = flag.String("p", "/home/john/tmp/spider/us_debt", "保存地址")
	fn   = flag.String("n", "2006_01_02", "保存文件名字")
	cn   = flag.String("c", fmt.Sprintf("日期： %s, 美国债务：", time.Now().Format("2006/1/2 15:04:05")), "保存内容")
	fg   = flag.Bool("f", true, "是否是数字")
	help = flag.Bool("h", false, "帮助")
)

const (
	helpStr = "/home/john/tmp/spider/spr -l https://tiyu.baidu.com/tokyoly/home/tab/%E5%A5%96%E7%89%8C%E6%A6%9C -x '//*[@id=\"sfr-app\"]/div/div[2]/div/div/div/main/section/div[1]/b-grouplist-sticky/div/div[3]/div/div/div[2]/div/div/a[1]/div[2]/div[1]' -p /home/john/tmp/spider/sport -n 2006_01_02 -c 中国奖牌榜: -f=false\n" +
		"-l 抓取链接\n" +
		"-x 抓取规则\n" +
		"-p 保存地址\n" +
		"-n 保存文件名字\n" +
		"-c 保存内容\n" +
		"-f 是否是数字\n" +
		"-h 是否是帮助"
)

func main() {
	flag.Parse()

	if *help {
		fmt.Println(helpStr)
		os.Exit(0)
	}

	fmName := time.Now().Format(*fn)
	fName := fmt.Sprintf("%s/%s", *fp, fmName)
	f, err := os.OpenFile(fName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		fm := fmt.Sprintf("打开文件%s失败： ", fName) + "%v \n"
		fmt.Fprintf(os.Stdout, fm, err)
		os.Exit(1)
	}

	resp, err := http.Get(*link)
	if err != nil {
		fm := fmt.Sprintf("抓取%s失败： ", *link) + "%v \n"
		fmt.Fprintf(f, fm, err)
		os.Exit(1)
	}
	if resp.StatusCode != 200 {
		fm := fmt.Sprintf("抓取%s失败， 状态码为： ", *link) + "%s \n"
		fmt.Fprintf(f, fm, resp.StatusCode)
		os.Exit(1)
	}
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(f, "解析出错： %v \n", err)
		os.Exit(1)
	}

	root, err := xmlpath.ParseHTML(bytes.NewReader(b))
	if err != nil {
		fmt.Fprintf(f, "xpath解析出错： %v \n", err)
		os.Exit(1)
	}

	path := xmlpath.MustCompile(*xp)
	if value, ok := path.String(root); ok {
		if *fg {
			preNum := strings.ReplaceAll(value, ",", "")
			preNum = strings.ReplaceAll(preNum, "$", "")

			if isNum(preNum) {
				splitIdx := 4
				out := ""
				for i, v := range splitNum(preNum) {
					out += getSplitNum(v, splitIdx)
					if i == 0 {
						out += "."
					}
				}
				fmt.Fprintf(f, "%s%s美元\r\n", *cn, out)
			} else {
				fmt.Fprintf(f, "%s%s\r\n", *cn, value)
			}
		} else {
			fmt.Fprintf(f, "%s%s\r\n", *cn, value)
		}

	}
}

func isNum(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func splitNum(s string) []string {
	return strings.Split(s, ".")
}

func getSplitNum(v string, splitIdx int) string {
	n := len(v)
	if n <= splitIdx {
		return v
	}
	return getSplitNum(v[:n-splitIdx], splitIdx) + "," + v[n-splitIdx:]
}
