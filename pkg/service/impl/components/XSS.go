package componets

import (
	dalfox "github.com/hahwul/dalfox/v2/lib"
	"log"
	"net/http"
)

func CheckXSS(response *http.Response) {
	mu.Lock()
	defer mu.Unlock()
	log.Println("start url xss scan ", response.Request.RequestURI)
	opt := dalfox.Options{
		Cookie:      response.Request.Header.Get("Cookie"),
		UserAgent:   response.Request.UserAgent(),
		WAFEvasion:  true,
		UseDeepDXSS: true,
	}
	target := dalfox.Target{
		URL:    response.Request.URL.String(),
		Method: response.Request.Method,
	}
	if len(response.Request.Cookies()) != 0 {
		target.Options = opt
	}
	result, err := dalfox.NewScan(target)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println(result)
	}

}
