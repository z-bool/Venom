package componets

import "sync"

var mu sync.Mutex

var BlackDomain = []string{"*.similarsites.*", "*.gov.cn"}
var BlackType = ""
