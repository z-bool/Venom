package componets

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func CheckWeakPassword(response *http.Response) {
	mu.Lock()
	defer mu.Unlock()
	request := response.Request
	reqtype := request.Header.Get("Content-Type")
	if strings.Contains(reqtype, "x-www-form-urlencoded") {
		reader := bufio.NewReader(request.Body)
		body, _ := io.ReadAll(reader)
		fmt.Println(string(body))
	}
}
