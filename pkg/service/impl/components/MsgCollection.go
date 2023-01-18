package componets

import (
	"github.com/z-bool/Venom/pkg/common/constract"
	"github.com/z-bool/Venom/pkg/entity"
	"github.com/z-bool/Venom/pkg/utils"
)

var Result = &entity.MsgCollection{
	Githubs:    make(map[string][]string),
	Telephones: make(map[string][]string),
	Emails:     make(map[string][]string),
	Persons:    make(map[string][]string),
}

// 收集Github地址
func collectGithubLink(githubList []string, responseBody []byte) []string {
	return utils.RemoveDuplicateElement(utils.AddToList(githubList, utils.RegxResult(constract.GITHUB_REGX, responseBody)))
}

// 收集手机号
func collectTelephone(telephoneList []string, responseBody []byte) []string {
	return utils.RemoveDuplicateElement(utils.AddToList(telephoneList, utils.RegxResult(constract.TELEPHONE_REGX, responseBody)))
}

// 收集邮箱
func collectEmail(emailList []string, responseBody []byte) []string {
	return utils.RemoveDuplicateElement(utils.AddToList(emailList, utils.RegxResult(constract.EMAIL_REGX, responseBody)))
}

// 收集身份证
func collectPerson(personList []string, responseBody []byte) []string {
	return utils.RemoveDuplicateElement(utils.AddToList(personList, utils.RegxResult(constract.PERSON_REGX, responseBody)))
}

// 总收集
func CollectMsg(reqUrl string, responseBody []byte) {
	mu.Lock()
	defer mu.Unlock()
	telephone := collectTelephone(Result.Telephones[reqUrl], responseBody)
	if len(telephone) != 0 {
		Result.Telephones[reqUrl] = telephone
	}
	github := collectGithubLink(Result.Githubs[reqUrl], responseBody)
	if len(github) != 0 {
		Result.Githubs[reqUrl] = github
	}
	person := collectPerson(Result.Persons[reqUrl], responseBody)
	if len(person) != 0 {
		Result.Persons[reqUrl] = person
	}
	email := collectEmail(Result.Emails[reqUrl], responseBody)
	if len(email) != 0 {
		Result.Emails[reqUrl] = email
	}
}
