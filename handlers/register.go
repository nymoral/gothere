package handlers

import (
	"gothere/database"
	"gothere/models"
	"gothere/password"
	"gothere/templates"
	"log"
	"net/http"
	"strings"
)

func RegisterGet(w http.ResponseWriter) {
	// /register GET method handler.
	// Just render's the form.

	templates.Render(w, "register", nil)
}

func RegisterPost(w http.ResponseWriter, r *http.Request) {
	// /register POST method handler.
	// Validates the form,
	// check's if username is availible,
	// and then creates a user and redirects to
	// /login .

	db := database.GetConnection()

	var user models.User
	// Model out of form data.
	user.Email = r.FormValue("email")
	user.Email = strings.ToLower(user.Email)
	user.Password = r.FormValue("password")
	user.Firstname = r.FormValue("firstname")
	if len(user.Firstname) > 0 {
		user.Firstname = strings.ToUpper(user.Firstname[0:1]) + strings.ToLower(user.Firstname[1:])
	}
	user.Lastname = r.FormValue("lastname")
	if len(user.Lastname) > 0 {
		user.Lastname = strings.ToUpper(user.Lastname[0:1]) + strings.ToLower(user.Lastname[1:])
	}

	repeat := r.FormValue("repeat")

	var old models.RegisterContext
	// Model for return form.
	// In case there the data wasn't valid
	old.Firstname = user.Firstname
	old.Lastname = user.Lastname
	old.Email = user.Email
	old.Flag = user.UserValidate(repeat)

	if old.Flag != "" {
		templates.Render(w, "register", old)
		return
	}

	pass, _ := database.GetPassword(db, user.Email)
	// Checks if user exists.
	if pass != "" {
		old.Flag = "Vartotojas su šiuo el. pašto adresu jau egzistuoja."
		templates.Render(w, "register", old)
		return
	}

	user.Password = password.NewPassword(user.Password)
	database.CreateUser(db, &user)
	// Creates a user in the db.
	http.Redirect(w, r, "/login", http.StatusFound)
	log.Printf("USER CREATED (%s)\n", user.Email)
}
