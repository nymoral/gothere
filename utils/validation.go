package utils

import (
    "regexp"
    "gothere/models"
)

const emailRegex = "[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*@(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?"
func emailValidation(mail string) (bool) {
    rp := regexp.MustCompile(emailRegex)
    return rp.MatchString(mail)
}

func UserValidate(user *models.User, repeat string) (bool) {
    /* 
     * Registration form validation.
     * Returns true/fales based on weather the form fits
     * requirements.
     */

    if user.Password != repeat {
        // Password don't math.
        return false
    }
    if len(repeat) < 6{
        // Password lenght.
        return false
    }
    if len(user.Firstname) < 1 || len(user.Lastname) < 1 || len(user.Firstname) > 20 || len(user.Lastname) > 30 {
        // To check if not empty and fits in the db.
        return false
    }

    if ! emailValidation(user.Email) || len(user.Email) > 50 {
        // To check if not empty and fits in the db.
        // TODO regex check if this is valid email address.
        return false
    }
    return true
}
