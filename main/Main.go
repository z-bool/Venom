package main

import (
	"bufio"
	"flag"
	Core "github.com/z-bool/Venom/pkg/service/impl"
	"github.com/z-bool/Venom/pkg/service/impl/Websocket"
	"io"
	"log"
	"net/http"
)

func init() {
	// 初始化根证书
	err := Core.NewCertificate().Init()
	if err != nil {
		log.Println("初始化根证书失败：" + err.Error())
		return
	}
}

func main() {
	port := flag.String("port", "9090", "listen port")
	nagle := flag.Bool("nagle", true, "connect remote use nagle algorithm")
	proxy := flag.String("proxy", "0", "tcp prxoy remote host")
	flag.Parse()
	if *port == "0" {
		log.Fatal("port required")
		return
	}
	// 启动服务
	s := Core.NewProxyServer(*port, *nagle, *proxy)
	// 注册http客户端请求事件函数
	s.OnHttpRequestEvent = func(request *http.Request) {
		// log.Println(request)
	}
	// 注册http服务器响应事件函数
	s.OnHttpResponseEvent = func(response *http.Response) {
		// contentType := response.Header.Get("Content-Type")
		var reader io.Reader
		// if strings.Contains(contentType, "json") {
		// 	reader = bufio.NewReader(response.Body)
		// 	if header := response.Header.Get("Content-Encoding"); header == "gzip" {
		// 		reader, _ = gzip.NewReader(response.Body)
		// 	}
		// 	body, _ := io.ReadAll(reader)
		// 	log.Println("HttpResponseEvent：" + string(body))
		// }
		reader = bufio.NewReader(response.Body)
		body, _ := io.ReadAll(reader)
		log.Println(string(body))
	}
	// 注册socket5服务器推送消息事件函数
	s.OnSocket5ResponseEvent = func(message []byte) {
		log.Println("Socket5ResponseEvent：" + string(message))
	}
	// 注册socket5客户端推送消息事件函数
	s.OnSocket5RequestEvent = func(message []byte) {
		log.Println("Socket5RequestEvent：" + string(message))
	}
	// 注册ws客户端推送消息事件函数
	s.OnWsRequestEvent = func(msgType int, message []byte, target *Websocket.Conn, resolve Core.ResolveWs) error {
		log.Println("WsRequestEvent：" + string(message))
		return target.WriteMessage(msgType, message)
	}
	// 注册w服务器推送消息事件函数
	s.OnWsResponseEvent = func(msgType int, message []byte, client *Websocket.Conn, resolve Core.ResolveWs) error {
		log.Println("WsResponseEvent：" + string(message))
		return resolve(msgType, message, client)
	}
	_ = s.Start()
}
