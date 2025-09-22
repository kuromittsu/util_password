package util_password

import (
	"errors"
	"strings"
)

type PasswordValidateInvalidList struct {
	invalidList []string
}

/*
	Get
*/
/*
	get raw invalid list
*/
func (p *PasswordValidateInvalidList) Get() []string {

	return p.invalidList
}

/*
	Get joined
*/
/*
	get joined invalid list by provided join string (default ", ")
	and also return error version
*/
func (p *PasswordValidateInvalidList) GetJoined(join string) (string, error) {

	if join == "" {
		join = ", "
	}
	output := strings.Join(p.invalidList, join)
	return output, errors.New(output)
}
