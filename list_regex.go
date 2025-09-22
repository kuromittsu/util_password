package util_password

var listRegexs map[validateKey]string = map[validateKey]string{
	MIN_8_LENGTH:  `^.{8,}$`,
	MIN_1_UPPER:   `[A-Z]`,
	MIN_1_LOWER:   `[a-z]`,
	MIN_1_NUMBER:  `\d`,
	MIN_1_SPECIAL: `[^a-zA-Z0-9]`,
}
