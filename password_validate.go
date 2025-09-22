package util_password

import "regexp"

func passwordValidate(text string, rules []validateKey, lang Lang) PasswordValidateResult {

	var invalidList []string

	for _, k := range rules {

		rgx := getRegex(k)
		msg := getMessageByLang(lang, k)

		if matched, _ := regexp.MatchString(rgx, text); !matched {
			invalidList = append(invalidList, msg)
		}
	}

	return newPasswordValidateResult(
		len(invalidList) == 0,
		PasswordValidateInvalidList{
			invalidList: invalidList,
		},
	)
}
