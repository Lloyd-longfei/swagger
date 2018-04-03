package main

import (
	"github.com/widuu/goini"
	"encoding/base64"
	"strings"
	"github.com/mikemintang/go-curl"
	"encoding/json"
)

var Couch struct {
	ch_host string
	ch_port string
	ch_user string
	ch_pass string
	ch_db   string
}

type chs struct {
	Rev string `json:"rev"`
}

type results struct {
	Seqs    string `json:"seq"`
	Id      string `json:"id"`
	Changes []chs  `json:"changes"`
	Deleted bool   `json:"deleted"`
}

type response struct {
	Result  []results `json:"results"`
	LastSeq string    `json:"last_seq"`
	Pending int       `json:"pending"`
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
	seq := "2-g1AAAAFReJzLYWBg4MhgTmHgzcvPy09JdcjLz8gvLskBCjMlMiTJ____PyuDOZEpFyjAbmGZZJxmaYGuGIf2JAUgmWQPMiGRAZcaB5CaePxqEkBq6vGqyWMBkgwNQAqobD4hdQsg6vYTUncAou4-IXUPIOpA7ssCAFGda4M"
	_, err := changes(seq)
	if err != nil {
		panic(err)
	}
	//res

}

func changes(seq string) (*[]interface{}, error) {
	var str = ""
	//couchdb link
	var opt = map[string]string{
		"limit": "2",
		"feed":  "longpoll",
		"seq":   seq,
	}
	for k, v := range opt {
		str += k + "=" + v + "&"
	}
	url := Couch.ch_host + ":" + Couch.ch_port + "/" + Couch.ch_db + "/_changes?" + strings.TrimRight(str, "&")
	body, err := curls(url)
	res := response{}
	json.Unmarshal([]byte(body), &res)
	var arrLists []interface{}
	Lists := make(map[string]string)
	for _, v := range res.Result {
		Lists["doc_id"] = v.Id
		Lists["seq"] = v.Seqs
		arrLists = append(arrLists, Lists)
	}
	return &arrLists, err
}

func docs(docid string) (string, error) {
	url := Couch.ch_host + ":" + Couch.ch_port + "/" + Couch.ch_db + "/" + docid
	body, err := curls(url)

	return body, err
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func curls(url string) (string, error) {
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

func redis() {

}
