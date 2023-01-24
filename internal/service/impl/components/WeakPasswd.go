package componets

import (
	"net/http"
)

func CheckWeakPassword(request *http.Request) {
	//var bodyStr string
	//
	//reader := bufio.NewReader(request.Body)
	//body, _ := io.ReadAll(reader)
	//bodyStr = string(body)
	//if utils.RegxRequest(constract.USER_NAME_REGX, string(body)) {
	//	username := utils.RegxFormData(constract.USER_NAME_DATA_REGX, constract.USER_NAME_REGX, string(body))
	//	// 将请求包中的参数替换，并且发起请求做出响应，记录URL
	//
	//}
}

// 将字典中的用户名和密码进行替换
func ReplaceUsernameAndPasswd(username string, usernameRegx string, passwd string, passwdRegx string, data string) {
	// 这里将作为一版本的临时处理
	//var usernameStr string
	//var passwdStr string
	//usernameKey := RegxResult(usernameRegx, []byte(data))
	//passwdKey := RegxResult(passwdRegx, []byte(data))
	//if len(usernameKey) != 0 {
	//	// 这里需要做个用户名遍历
	//
	//	usernameStr = regexp.MustCompile(usernameRegx).ReplaceAllString(data, strings.Replace(usernameKey[0], username, "123", 1))
	//}

}
