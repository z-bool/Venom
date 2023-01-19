package componets

import (
	"github.com/z-bool/Venom/pkg/common/constract"
	"github.com/z-bool/Venom/pkg/entity"
	"github.com/z-bool/Venom/pkg/utils"
	"net/http"
)

var CorsUrl = make([]entity.URL, 0)

func Cors(res *http.Response, responseBody []byte) {
	mu.Lock()
	defer mu.Unlock()
	req := res.Request
	var title string
	headers := ""
	titles := utils.RegxResult(constract.TitleRegx, responseBody)
	if len(titles) == 0 {
		title = ""
	} else if len(titles) > 1 {
		title = titles[len(titles)-1]
	}
	headerObj := res.Header
	headerObj.Del("Cookie")
	for key, headerOption := range headerObj {
		val := ""
		for _, header := range headerOption {
			val += header
		}
		headers += key + ": " + val + "\n"
	}
	if res.Header.Get(constract.HEADER_ACCESS_CONTROL_ALLOW_CREDENTIALS) == "true" && res.Header.Get(constract.HEADER_ACCESS_CONTROL_ALLOW_ORIGIN) != "" {
		uri := entity.URL{
			Domain:          res.Request.Host,
			Url:             req.URL.Path,
			Title:           title,
			ResponseHeaders: headers,
		}
		CorsUrl = append(CorsUrl, uri)
	}
}
