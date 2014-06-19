package utils

import (
	"fmt"
	"regexp"
)

const (
	emailRegex  = "[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*@(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?"
	resultRegex = `^\s*(\d+)\s*[:|-]\s*(\d+)\s*$`
)

var (
	emailR  *regexp.Regexp
	resultR *regexp.Regexp
)

func init() {
	emailR = regexp.MustCompile(emailRegex)
	resultR = regexp.MustCompile(resultRegex)
}

func EmailValidation(mail string) bool {
	return emailR.MatchString(mail)
}

func ExtractResult(s string, f1, f2 *int) bool {
	sl := resultR.FindStringSubmatch(s)
	if len(sl) != 3 {
		return false
	} else {
		fmt.Sscanf(sl[1], "%d", f1)
		fmt.Sscanf(sl[2], "%d", f2)
		return true
	}

}
