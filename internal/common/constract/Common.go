package constract

const TitleRegx = "<title>(.*?)</title>"

// 请求响应过滤
var BlackDomain = []string{"*.similarsites.*", "*.gov.cn", "*.shodan.io"}

const BlackXSSType = "(https|http)://([^\\s]*)/([^\\s]*)\\.(js|css|jpg|jpeg|png|webp|mp4|mp3|ttf|woff|woff2|word|txt|xlxs|avif|csv|ico)(\\?)?([^\\s]*)"
const USER_NAME_REGX = "&?(user|username|account|User|UserName|Username|Account)="
const USER_ID_REGX = ""
const USER_NAME_DATA_REGX = "(user|username|account)=([0-9a-zA-Z_-]*)?"
const USER_NAME_FULL_REGX = "(user|username|account|User|UserName|Username|Account)=[0-9a-zA-Z_]*"
