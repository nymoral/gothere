package handlers

import (
	"database/sql"
	"fmt"
	"gothere/database"
	"gothere/password"
	"gothere/templates"
	"gothere/utils"
	"log"
	"net/http"
)

func ForgotGet(w http.ResponseWriter, r *http.Request) {
	forgotKey := r.URL.RawQuery

	db := database.GetConnection()

	if forgotKey == "" {
		// A simple GET for the form.
		templates.Render(w, "forgot", nil)
	} else if database.CheckRecovery(db, forgotKey) {
		// Render recover page.
		templates.Render(w, "recover", forgotKey)
	} else {
		// Redirect to login.
		http.Redirect(w, r, "/login/", http.StatusFound)
	}
}

func recovery(db *sql.DB, msg utils.Message, pk int, key string) {
	err := msg.Send()
	if err != nil {
		log.Println(err)
	} else {
		// Create a record only if the email succeeds.
		database.CreateRecovery(db, key, pk)
	}
}

func ForgotPost(w http.ResponseWriter, r *http.Request) {
	forgotKey := r.URL.RawQuery
	db := database.GetConnection()
	if forgotKey == "" {
		// Initial forgot submit.
		email := r.FormValue("email")
		pk, _ := database.GetPkAdmin(db, email)
		if pk != -1 {
			if !database.RecoveryExists(db, pk) {
				// Send out an email.
				key := utils.GenRecoveryKey()
				msg := utils.Message{email,
					"Slaptažodio atkūrimas",
					fmt.Sprintf("Norėdami atkurti slaptažodį eikite į:\n\nhttp://futbolas.aivaras.in/forgot/?%s\n\nŠi nuoroda galios dvi dienas.", key)}
				go recovery(db, msg, pk, key)
				log.Printf("An email was sent to %s\n", email)
				templates.Render(w, "forgot", nil)
			} else {
				// The recovery already exists.
				templates.Render(w, "forgot", nil)
			}
		} else {
			//  User non existing.
			templates.Render(w, "forgot", nil)
		}
	} else {
		// Actual recovery.
		newPassword := r.FormValue("new")
		repeat := r.FormValue("repeat")
		if len(newPassword) < 6 || newPassword != repeat {
			// Bad password
			templates.Render(w, "recover", forgotKey)
		} else {
			hashed := password.NewPassword(newPassword)
			database.DoRecover(db, forgotKey, hashed)
			http.Redirect(w, r, "/login", http.StatusFound)
			log.Printf("KEY (%s) RESET PASSWORD\n", forgotKey)
		}
	}
}
