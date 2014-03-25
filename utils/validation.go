package utils

import (
    "regexp"
)

const emailRegex = "[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*@(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?"

func EmailValidation(mail string) (bool) {
    rp := regexp.MustCompile(emailRegex)
    return rp.MatchString(mail)
}

