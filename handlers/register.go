package handlers

import (
    "gothere/models"
    "net/http"
    "gothere/templates"
    "gothere/database"
    "gothere/utils"
    "gothere/password"
    "database/sql"
)

func RegisterGet(w http.ResponseWriter) {
    templates.Render(w, "register", nil)
}

func RegisterPost(w http.ResponseWriter, r *http.Request, db *sql.DB) {
    email := r.FormValue("email")
    pass := r.FormValue("password")
    repeat := r.FormValue("repeat")
    firstname := r.FormValue("firstname")
    lastname := r.FormValue("lastname")

    var user models.User
    user.Email = email
    user.Password = pass
    user.Firstname = firstname
    user.Lastname = lastname
    if ! utils.UserValidate(user, repeat) {
        templates.Render(w, "register", "Form Not valid")
    } else if database.GetPassword(db, email) != "" {
        templates.Render(w, "register", "User already exists")
    } else {
        user.Password = password.NewPassword(pass)
        database.CreateUser(db, user)
        http.Redirect(w, r, "/login", 302) 
    }
}


