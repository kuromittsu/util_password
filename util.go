package util_password

func getMessageByLang(lang Lang, key validateKey) string {
	value := listMessageByLangs[lang][key]
	return value
}

func getRegex(key validateKey) string {
	value := listRegexs[key]
	return value
}
