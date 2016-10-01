package common

import "regexp"

func CheckEmail(checkedString string) bool {
	rs, err := regexp.MatchString(EMAIL_REGX, checkedString)
	Check(err)

	return rs
}

func CheckPhoneNumber(numberPhone string, countryId string) {

}
