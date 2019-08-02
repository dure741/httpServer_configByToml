package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Base struct {
		Ipaddr     string `toml:"ipaddr"`
		Port       string `toml:"port"`
		FileServer string `toml:"fileServer"`
		Host       string `toml:"host"`
		Route      string `toml:"route"`
	} `toml:"Base"`
	Wait struct {
		Time time.Duration `toml:"timeWaitForDebug"`
	} `toml:"Wait"`
}

var config Config

// hello world, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
	//w,用来给客户端回复数据
	//req,用来读取客户端发送的数据
	fmt.Println("config waiting time : ", config.Wait.Time)
	time.Sleep(config.Wait.Time)
	io.WriteString(w, "hello, world!\n")
}

func main() {
	//读取toml配置文件
	if _, err := toml.DecodeFile("./config.toml", &config); err != nil {
		log.Fatal(err)
		fmt.Println("error : ", err)
		return
	}

	//注册处理函数，用户连接，自动调用指定的处理函数
	go http.HandleFunc(config.Base.Route, HelloServer) //回调函数
	//go http.HandleFunc("www.a.com/", HelloServer)
	http.Handle(config.Base.FileServer, http.StripPrefix(config.Base.FileServer, http.FileServer(http.Dir( /*"/root/secondWeek/httpSever_GoModule/httpserver"*/ "./"+config.Base.FileServer))))

	//监听绑定(地址:端口)
	err := http.ListenAndServe(config.Base.Ipaddr+":"+config.Base.Port, nil)
	if err != nil {
		log.Fatal("ListenAndServe IP and Port error : ", err)
		fmt.Println("ListenAndServe IP and Port error : ", err)
	}

	//监听绑定(域名)
	/*
		err1 := http.ListenAndServe(config.Base.Host+":"+config.Base.Port, nil)
		if err1 != nil {
			log.Fatal("ListenAndServe Host error : ", err1)
			fmt.Println("ListenAndServe Host error : ", err1)
		}
	*/
}
