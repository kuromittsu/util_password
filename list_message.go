package util_password

var listMessageByLangs map[Lang]map[validateKey]string = map[Lang]map[validateKey]string{
	LANG_ID: {
		MIN_8_LENGTH:  "minimal 8 karakter",
		MIN_1_UPPER:   "minimal 1 huruf besar",
		MIN_1_LOWER:   "minimal 1 huruf kecil",
		MIN_1_NUMBER:  "minimal 1 angka (0-9)",
		MIN_1_SPECIAL: "minimal 1 karakter spesial",
	},
	LANG_EN: {
		MIN_8_LENGTH:  "at least 8 characters",
		MIN_1_UPPER:   "at least 1 uppercase letter",
		MIN_1_LOWER:   "at least 1 lowercase letter",
		MIN_1_NUMBER:  "at least 1 number (0-9)",
		MIN_1_SPECIAL: "at least 1 special character",
	},
}
