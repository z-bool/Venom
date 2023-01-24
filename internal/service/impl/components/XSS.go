package componets

import (
	dalfox "github.com/hahwul/dalfox/v2/lib"
	"github.com/z-bool/Venom/internal/common/constract"
	"github.com/z-bool/Venom/internal/utils"
	"log"
	"net/http"
)

func CheckXSS(response *http.Response) {
	if !utils.RegxRequest(constract.BlackXSSType, response.Request.URL.String()) {
		log.Println("start url xss scan ", response.Request.RequestURI)
		var cookie string
		if len(response.Request.Cookies()) != 0 {
			cookie = response.Request.Header.Get("Cookie")
		}
		opt := dalfox.Options{
			Cookie:      cookie,
			UserAgent:   response.Request.UserAgent(),
			WAFEvasion:  true,
			UseDeepDXSS: true,
		}
		target := dalfox.Target{
			URL:     response.Request.URL.String(),
			Method:  response.Request.Method,
			Options: opt,
		}

		result, err := dalfox.NewScan(target)
		if err != nil {
			log.Println(err)
		} else {
			log.Println(result)
		}
	}
}
