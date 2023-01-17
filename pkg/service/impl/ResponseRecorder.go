package Core

import (
	componets "github.com/z-bool/Venom/pkg/service/impl/components"
)

func Collect(reqUrl string, responseBody []byte) {
	go componets.CollectMsg(reqUrl, responseBody)
}
