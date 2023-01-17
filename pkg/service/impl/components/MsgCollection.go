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
	Result.Telephones[reqUrl] = collectTelephone(Result.Telephones[reqUrl], responseBody)
	Result.Githubs[reqUrl] = collectGithubLink(Result.Githubs[reqUrl], responseBody)
	Result.Persons[reqUrl] = collectPerson(Result.Persons[reqUrl], responseBody)
	Result.Emails[reqUrl] = collectEmail(Result.Emails[reqUrl], responseBody)
}
