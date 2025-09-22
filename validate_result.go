package util_password

func newPasswordValidateResult(isValid bool, invalidList PasswordValidateInvalidList) PasswordValidateResult {
	return PasswordValidateResult{
		isValid:     isValid,
		invalidList: invalidList,
	}
}

type PasswordValidateResult struct {
	isValid     bool
	invalidList PasswordValidateInvalidList
}

/*
	Is valid
*/
/*
	if valid return true and false otherwise
*/
func (p *PasswordValidateResult) IsValid() bool {
	return p.isValid
}

/*
	Invalid list
*/
/*
	return PasswordValidateInvalidList struct
*/
func (p *PasswordValidateResult) InvalidList() *PasswordValidateInvalidList {
	return &p.invalidList
}
