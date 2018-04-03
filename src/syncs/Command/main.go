package main

import (
	"encoding/json"
	"fmt"
)

type Server struct {
	ServerName string
	ServerIP   string
}

type Serverslice struct {
	Sev string
	Servers []Server
}

func main() {
	var s Serverslice
	str := `{"sev":"dsaads","servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
	json.Unmarshal([]byte(str), &s)
	for _,v := range s.Servers{
		fmt.Println(v)
	}
}