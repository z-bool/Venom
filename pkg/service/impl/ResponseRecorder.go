package Core

import (
	"bufio"
	componets "github.com/z-bool/Venom/pkg/service/impl/components"
	"io"
	"net/http"
)

// 记录爬虫URL的记录，防止爬取重复
var Urls = make([]string, 0)

func Collect(response *http.Response) {
	//reqUrl := response.Request.URL.RequestURI()
	var reader io.Reader
	reader = bufio.NewReader(response.Body)
	body, _ := io.ReadAll(reader)
	componets.CollectMsg(response.Request.Host, body)
	componets.Cors(response, body)
	Print()
}

func Print() {
	//fmt.Println(componets.Result)
	//fmt.Println(componets.CorsUrl)
}
