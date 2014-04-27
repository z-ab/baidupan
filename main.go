package main

import (
	"encoding/json"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/lunny/xorm"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var engine *xorm.Engine
var dberr error
var ch chan string

var ch1 chan int

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	engine, dberr = xorm.NewEngine("mysql", "root:root@tcp(127.0.0.1:3306)/res?charset=utf8")
	//engine.ShowSQL = true
	if dberr != nil {
		fmt.Println(dberr)
		return
	}
	//defer engine.Close()
}
func main() {
	do := flag.String("do", "uk", "命令默认uk（获取百度用户的uk），其他命令ip（获取代理ip）,res(获取百度共享资源)")
	flag.Parse()
	switch *do {
	case "uk":
		ReqFollow()
	case "ip":
		CrawlIpUrl()
	case "res":
		ObRes()
	default:
		fmt.Println("command is not right ,use -h or use -help")
	}

	////for {
	//ReqFollow() //获取用户
	//time.Sleep(time.Second * 5)
	////}

	////Veriip()
	//Obuk()
	//Crawresstart()
}

func ObRes() {
	sql := "select * from buk where `state`=1 and share>100 order by `share`  "
	result, err := engine.Query(sql)
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(result) < 1 {
		return
	}
	ch = make(chan string, 20)
	for _, v := range result {
		url := "http://yun.baidu.com/pcloud/feed/getsharelist?t=1396440785121&category=0&auth_type=1&start=0&limit=60&query_uk=" + string(v["uk"])
		go Crawresstart(url)

		<-ch

		engine.Exec("update buk set `state`=2 where `uk`=?", v["uk"])
	}

}
func Crawresstart(url string) {

	total := Crawlres(url)
	fmt.Println(total)
	if total > 0 {
		for i := 1; i < (total/60)+1; i++ {
			t := Crawlres(strings.Replace(url, "start=0", "start="+strconv.Itoa(60*i), 1))
			fmt.Println(t)
		}
	}
	ch <- url
}
func Crawlres(url string) int {
	fmt.Println(url)
	b, err := curl(url)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	var res Res
	json.Unmarshal(b, &res)
	if res.Total_count == 0 {
		return 0
	}
	sql := "insert  into `res`(`uk`,`shareid`,`shorturl`,`category`,`title`,`username`,`dateid`,`file_name`,`size`,`path`,`md5`,`thumburl`,`sign`,`t_stamp`,`fid`,`addtime`,`updatetime`,`source_uid`,`source_id`,`vcnt`,`dcnt`,`tcnt`,`like_status`,`like_count`,`comment_count`) values"
	sql += "(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	for _, v := range res.Records {
		for _, v1 := range v.Filelist {
			if v.Category != 2 && v.Category != 3 {
				_, e := engine.Exec(sql, v.Uk, v.Shareid, v.Shorturl, v.Category, v.Title, v.Username, v.Data_id, v1.Server_filename, v1.Size, v1.Path, v1.Md5, v1.Thumburl, v1.Sign, v1.Time_stamp, v1.Fs_id, time.Now().Unix(), time.Now().Unix(), v.Source_uid, v.Source_id, v.VCnt, v.DCnt, v.TCnt, v.Like_status, v.Like_count, v.Comment_count)
				if e != nil {
					fmt.Println(e)
					time.Sleep(time.Second * 3)
				}
			}
		}
	}
	return res.Total_count
}
