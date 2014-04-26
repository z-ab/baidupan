package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func ReqFollow() {
	sql := "select * from buk where `state`=0 "
	result, err := engine.Query(sql)
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(result) < 1 {
		return
	}
	for _, v := range result {
		url := "http://yun.baidu.com/pcloud/friend/getfollowlist?query_uk=" + string(v["uk"]) + "&limit=24&start=0"
		ReqFollowUrl(url)
		engine.Exec("update buk set `state`=1 where `uk`=?", v["uk"])
	}
}
func ReqFollowUrl(url string) {
	total := Follow(url)
	fmt.Println(total)
	if total > 0 {
		for i := 1; i < (total/24)+1; i++ {
			t := Follow(strings.Replace(url, "start=0", "start="+strconv.Itoa(24*i), 1))
			fmt.Println(t)
		}
	}
}
func Follow(url string) int {
	fmt.Println(url)
	b, e := curl(url)
	if e != nil {
		fmt.Println(e)
		return -1
	}
	var fr Followres
	json.Unmarshal(b, &fr)
	if fr.Total_count == 0 {
		return 0
	}
	sql := "insert ignore into buk(`name`,`uk`,`fans`,`album`,`follow`,`share`,`updatetime`) values(?,?,?,?,?,?,?)"
	for _, v := range fr.Followlist {
		if v.Follow_count != 0 || v.Pubshare_count != 0 {
			_, e := engine.Exec(sql, v.Follow_uname, v.Follow_uk, v.Fans_count, v.Album_count, v.Follow_count, v.Pubshare_count, time.Now().Unix())
			if e != nil {
				fmt.Println(e)
				continue
			}
		}
	}
	return fr.Total_count
}
