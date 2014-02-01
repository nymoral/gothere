package utils

import (
    "gothere/models"
)

func UserValidate(user models.User, repeat string) (bool) {
    if user.Password != repeat || len(repeat) < 6{
        return false
    }
    if len(user.Firstname) < 2 || len(user.Lastname) < 2 {
        return false
    }

    if len(user.Email) < 5 {
        return false
    }
    return true
}
