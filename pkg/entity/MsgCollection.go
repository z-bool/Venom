package entity

type MsgCollection struct {
	Telephones map[string][]string
	Emails     map[string][]string
	Persons    map[string][]string
	Githubs    map[string][]string
}
