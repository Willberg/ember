package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/glog"
	"github.com/jmoiron/sqlx"
	"gopkg.in/xmlpath.v2"
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
	help = flag.Bool("h", false, "帮助")
)

const (
	helpStr = "/home/john/tmp/spider/spr -l https://tiyu.baidu.com/tokyoly/home/tab/%E5%A5%96%E7%89%8C%E6%A6%9C -x '//*[@id=\"sfr-app\"]/div/div[2]/div/div/div/main/section/div[1]/b-grouplist-sticky/div/div[3]/div/div/div[2]/div/div/a[1]/div[2]/div[1]'  -log_dir='/home/john/tmp/spider/log'\n" +
		"-l 抓取链接\n" +
		"-x 抓取规则\n" +
		"-log_dir 错误日志目录\n"
)

type Debt struct {
	Id         int     `db:"id"`
	Date       string  `db:"date"`
	Debt       float64 `db:"debt"`
	CreateTime int     `db:"create_time"`
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if *help {
		fmt.Println(helpStr)
		os.Exit(0)
	}

	b, err := fetch(*link)
	if err != nil {
		glog.Error(err)
		os.Exit(1)
	}

	debt, err := parse(b)
	if err != nil {
		glog.Error(err)
		os.Exit(1)
	}

	save(debt)
}

func parse(b []byte) (float64, error) {
	root, err := xmlpath.ParseHTML(bytes.NewReader(b))
	if err != nil {
		errMsg := fmt.Sprintf("xpath解析出错： %v \n", err)
		return 0, errors.New(errMsg)
	}

	path := xmlpath.MustCompile(*xp)
	if value, ok := path.String(root); ok {
		debt := strings.ReplaceAll(value, ",", "")
		debt = strings.ReplaceAll(debt, "$", "")

		if isNum(debt) {
			return strconv.ParseFloat(debt, 64)
		} else {
			return 0, errors.New("非数字")
		}
	}
	return 0, errors.New("xpath解析错误")
}

func save(debt float64) {
	db, err := sqlx.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/g_us")
	if err != nil {
		glog.Error("open mysql failed, %v", err)
		return
	}
	defer db.Close()

	//查询是否已经入库
	var debts []Debt
	date := time.Now().Format("2006-01-02")
	err = db.Select(&debts, "select id, `date`, debt, create_time from us_debt where `date`=?", date)
	if err != nil {
		glog.Error("select failed, %v", err)
		return
	}

	if len(debts) == 0 {
		r, err := db.Exec("insert into us_debt(`date`, debt, create_time) values(?, ?, ?)", date, debt, time.Now().Unix())
		if err != nil {
			glog.Error("exec failed, %v", err)
			return
		}
		id, err := r.LastInsertId()
		if err != nil {
			glog.Error("exec failed, ", err)
			return
		}

		fmt.Println("insert success:", id)
	}
}

func fetch(link string) ([]byte, error) {
	resp, err := http.Get(link)
	if err != nil {
		f := fmt.Sprintf("抓取%s失败： ", link) + "%v \n"
		errMsg := fmt.Sprintf(f, err)
		return nil, errors.New(errMsg)
	}
	if resp.StatusCode != 200 {
		f := fmt.Sprintf("抓取%s失败， 状态码为： ", link) + "%s \n"
		errMsg := fmt.Sprintf(f, err)
		return nil, errors.New(errMsg)
	}
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		errMsg := fmt.Sprintf("解析出错： %v \n", err)
		return nil, errors.New(errMsg)
	}
	return b, nil
}

func isNum(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}
