package Core

import (
	componets "github.com/z-bool/Venom/internal/service/impl/components"
	"net/http"
)

func HookRequest(request *http.Request) {
	componets.CheckWeakPassword(request)
}
