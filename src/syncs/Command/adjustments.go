package main

import (
	"sync/Database"
	"log"
	"github.com/widuu/goini"
)

type option struct {

	tablename string
	dbname    string
	host      string
	port      string
	user      string
	password  string
	sslmode   string
	schema    string
}

var Opt option

func main() {
	Opt.tablename = "adjustments"
	conf := goini.SetConfig("../env.ini")
	Opt.dbname = conf.GetValue("original", "dbname")
	Opt.host = conf.GetValue("original", "host")
	Opt.port = conf.GetValue("original", "port")
	Opt.user = conf.GetValue("original", "user")
	Opt.password = conf.GetValue("original", "password")
	Opt.sslmode = conf.GetValue("original", "sslmode")
	Opt.schema = conf.GetValue("original", "schema")
	cfg := "dbname=" + Opt.dbname + " host=" + Opt.host + " port=" + Opt.port + " user=" + Opt.user + " password=" + Opt.password + " sslmode=" + Opt.sslmode
	res := Database.Exec(cfg, createTable(Opt.tablename, Opt.schema))
	log.Println(res)

}

func createTable(table string, schema string) string {

	sql := "create table if not exists " + schema + "." + table + `
		(
			"uuid" uuid NOT NULL DEFAULT uuid_generate_v1(),
			"type" varchar NOT NULL COLLATE "default",
			"label" varchar NOT NULL COLLATE "default",
			"number" numeric NOT NULL,
			"product_uuid" uuid,
			"order_uuid" uuid,
			"currency_code" char(3) NOT NULL COLLATE "default",
			"included" bool)WITH (OIDS=FALSE);
	`
	return sql
}
