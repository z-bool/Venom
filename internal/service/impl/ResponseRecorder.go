package Core

import (
	"net/http"
)

// 记录爬虫URL的记录，防止爬取重复
var Urls = make([]string, 0)

// 首次登录记录该域名的访问Cookie
var BeforeCookie = make(map[string]string, 0)

// 后续Cookie存在更新都在此替换，用于权限校验
var AfterCookie = make(map[string]string, 0)

//var BlackDomain = [""]

func Collect(response *http.Response) {
	//reqUrl := response.Request.URL.RequestURI()
	//var reader io.Reader
	//reader = bufio.NewReader(response.Body)
	//body, _ := io.ReadAll(reader)
	//go componets.CollectMsg(response.Request.Host, body)
	//componets.Cors(response, body)
	//go componets.CheckXSS(response)
	//Print()
}

func Print() {
	//fmt.Println(componets.Result)
	//fmt.Println(componets.CorsUrl)
}
