package util_password

func PasswordHash(rawPassword string) (string, error) {
	return passwordHash(rawPassword)
}

func PasswordCompare(hashedPassword, rawPassword string) error {
	return passwordCompare(hashedPassword, rawPassword)
}

func PasswordValidate(text string, rules []validateKey, lang Lang) PasswordValidateResult {
	return passwordValidate(text, rules, lang)
}

func DefaultRules() []validateKey {
	return []validateKey{
		MIN_8_LENGTH,
		MIN_1_UPPER,
		MIN_1_LOWER,
		MIN_1_NUMBER,
		MIN_1_SPECIAL,
	}
}
