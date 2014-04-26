package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	_ "github.com/go-sql-driver/mysql"
	"net"
	"strconv"
	"strings"
	"time"
)

func Veriip() {
	//Crawlurl()

	sql := "select * from ip where `status`=1 and updatetime<? "
	result, err := engine.Query(sql, time.Now().Unix()-600)
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(result) < 1 {
		return
	}
	for k, v := range result {
		fmt.Println(k, string(v["ip"]))
		//_, err := net.Dial("tcp", string(v["ip"]))

		func() {
			start := time.Now().UnixNano()
			_, err := net.DialTimeout("tcp", string(v["ip"]), time.Second*3)
			end := time.Now().UnixNano()
			s := (end - start) / 1000 / 1000
			fmt.Println(s)
			if err != nil {
				engine.Exec("update `ip` set status=2 , updatetime=?, consume=? where `ip`=?", time.Now().Unix(), s, string(v["ip"]))
			} else {
				engine.Exec("update `ip` set  updatetime=? , consume=? where `ip`=?", time.Now().Unix(), s, string(v["ip"]))
			}

		}()

	}
}

func Crawlurl() {
	ipurl := ""
	id := "1873"
	for i := 0; i < 11; i++ {
		if i == 0 {
			ipurl = "http://www.youdaili.cn/Daili/http/" + id + ".html"
		} else {
			ipurl = "http://www.youdaili.cn/Daili/http/" + id + "_" + strconv.Itoa(i) + ".html"
		}
		fmt.Println(ipurl)

		Crawlip(ipurl)
	}
}
func Crawlip(ipurl string) {
	//engine, err := xorm.NewEngine("mysql", "root:root@tcp(10.33.3.222:3306)/res?charset=utf8")
	////engine.ShowSQL = true
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//defer engine.Close()
	var doc *goquery.Document
	var e error

	if doc, e = goquery.NewDocument(ipurl); e != nil {
		fmt.Println(e)
		return
	}
	doc.Find(".cont_font p").Each(func(idx int, s *goquery.Selection) {
		iptxt := s.Text()
		iparr := strings.Split(iptxt, "\n")
		if len(iparr) < 1 {
			return
		}
		sql := "insert ignore into ip(`ip`,`area`,`addtime`) values(?,?,?)"
		for _, ipstr := range iparr {
			fmt.Println(ipstr)
			//ip := strings.SplitN(ipstr, ".", 2)
			if ipstr != "" {
				retip := strings.Split(strings.Replace(strings.Replace(ipstr, "@", "||", -1), "#", "||", -1), "||")
				fmt.Println(retip)
				//break
				engine.Exec(sql, retip[0], retip[2], time.Now().Unix())

			}
		}
	})
}
