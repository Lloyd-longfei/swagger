package main

import (
	"github.com/widuu/goini"
	"encoding/base64"
	"log"
	"strings"
	"github.com/mikemintang/go-curl"
)

var Couch struct {
	ch_host string
	ch_port string
	ch_user string
	ch_pass string
	ch_db   string
}

func init() {
	conf := goini.SetConfig("../env.ini")
	/*
		couch link
	 */
	Couch.ch_host = conf.GetValue("couchdb", "host")
	Couch.ch_port = conf.GetValue("couchdb", "port")
	Couch.ch_user = conf.GetValue("couchdb", "user")
	Couch.ch_pass = conf.GetValue("couchdb", "pass")
	Couch.ch_db = conf.GetValue("couchdb", "dbname")
}

func main() {

	//couchdb link
	var opt = map[string]string{
		"limit": "1",
		"feed":  "longpoll",
		"seq":   "2-g1AAAAFReJzLYWBg4MhgTmHgzcvPy09JdcjLz8gvLskBCjMlMiTJ____PyuDOZEpFyjAbmGZZJxmaYGuGIf2JAUgmWQPMiGRAZcaB5CaePxqEkBq6vGqyWMBkgwNQAqobD4hdQsg6vYTUncAou4-IXUPIOpA7ssCAFGda4M",
	}
	res, err := couch(opt)
	if err != nil {
		panic(err)
	}
	log.Println(res)

}

func changes(opt map[string]string)  {
	var str = ""

	for k, v := range opt {
		str += k + "=" + v + "&"
	}
	url := Couch.ch_host + ":" + Couch.ch_port + "/" + Couch.ch_db + "/_changes?" + strings.TrimRight(str, "&")

}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func curls(url string) (string, error){
	headers := map[string]string{
		"Cache-Control": "no-cache",
		"Authorization": "Basic " + basicAuth(Couch.ch_user, Couch.ch_pass),
		"Content-Type":  "application/json",
	}
	req := curl.NewRequest()
	resp, err := req.
		SetUrl(url).
		SetHeaders(headers).
		Get()

	return resp.Body, err
}

func seq(){

}

func redis() {

}
