package Vendor

import (
	"os"
	"fmt"
	"log"
	"time"
)


func ReadLineLog(errs error) {
	var filename = "./Log/"+time.Now().Format("2006-01") +".log"
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0666) //打开文件
	if err != nil {
		fmt.Printf("%s\r\n", err.Error())
		os.Exit(-1)
	}
	defer f.Close()
	logger := log.New(f, "\r\n", log.Ldate|log.Ltime|log.Llongfile)
	logger.Println(errs.Error())
}
