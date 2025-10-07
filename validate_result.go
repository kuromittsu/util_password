package util_password

func newPasswordValidateResult(isValid bool, invalidList passwordValidateInvalidList) passwordValidateResult {
	return passwordValidateResult{
		isValid:     isValid,
		invalidList: invalidList,
	}
}

type passwordValidateResult struct {
	isValid     bool
	invalidList passwordValidateInvalidList
}

/*
	Is valid
*/
/*
	if valid return true and false otherwise
*/
func (p *passwordValidateResult) IsValid() bool {
	return p.isValid
}

/*
	Invalid list
*/
/*
	return passwordValidateInvalidList struct
*/
func (p *passwordValidateResult) InvalidList() *passwordValidateInvalidList {
	return &p.invalidList
}
