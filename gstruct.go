package main

type Res struct {
	Error       int       `json:"error"`
	Total_count int       `json:"total_count"`
	Records     []Records `json:"records"`
}
type Records struct {
	Category       int        `json:"category"`
	Title          string     `json:"title"`
	Shareid        string     `json:"shareid"`
	Data_id        string     `json:"data_id"`
	Filecount      int64      `json:"filecount"`
	Uk             int64      `json:"uk"`
	Username       string     `json:"username"`
	Feed_time      int64      `json:"feed_time"`
	Desc           string     `json:"desc"`
	Avatar_url     string     `json:"avatar_url"`
	Category_1_cnt int64      `json:"category_1_cnt"`
	Filelist       []Filelist `json:"filelist"`
	Shorturl       string     `json:"shorturl"`
	Source_uid     int64      `json:"source_uid"`
	Source_id      int64      `json:"source_id"`
	VCnt           int64      `json:"vCnt"`
	DCnt           int64      `json:"dCnt"`
	TCnt           int64      `json:"tCnt"`
	Like_status    int64      `json:"like_status"`
	Like_count     int64      `json:"like_count"`
	Comment_count  int64      `json:"comment_count"`
}
type Filelist struct {
	Server_filename string `json:"server_filename"`
	Size            int64  `json:"size"`
	Fs_id           int64  `json:"fs_id"`
	Path            string `json:"path"`
	Md5             string `json:"md5"`
	Thumburl        string `json:"thumburl"`
	Sign            string `json:"sign"`
	Time_stamp      int64  `json:"time_stamp"`
}
type Followres struct {
	Error       int          `json:"errno"`
	Total_count int          `json:"total_count"`
	Followlist  []Followlist `json:"follow_list"`
}
type Followlist struct {
	Album_count    int64  `json:"album_count"`
	Avatar_url     string `json:"avatar_url"`
	Fans_count     int64  `json:"fans_count"`
	Follow_count   int64  `json:"follow_count"`
	Follow_time    int64  `json:"follow_time"`
	Follow_uk      int64  `json:"follow_uk"`
	Follow_uname   string `json:"follow_uname"`
	Intro          string `json:"intro"`
	Is_vip         int64  `json:"is_vip"`
	Pubshare_count int64  `json:"pubshare_count"`
	Types          int64  `json:"types"`
	User_type      int64  `json:"user_type"`
}
