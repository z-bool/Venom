package constract

const GITHUB_REGX = "(https|http|ssh|git)(://|@)([^\\s]*?)\\.git|(https|http)://git(hub|lib|ee).com/[^\\s]*?/[a-zA-Z0-9_-]*"

const TELEPHONE_REGX = "1(3\\d|4[5-9]|5[0-35-9]|6[2567]|7[0-8]|8\\d|9[0-35-9])\\d{8}"
const EMAIL_REGX = "\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*"
const PERSON_REGX = "(([1-6][1-9]|50)\\d{4}(18|19|20)\\d{2}((0[1-9])|10|11|12)(([0-2][1-9])|10|20|30|31)\\d{3}[0-9Xx])|(([1-6][1-9]|50)\\d{4}\\d{2}((0[1-9])|10|11|12)(([0-2][1-9])|10|20|30|31)\\d{3}$)"
