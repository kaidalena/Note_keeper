package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"note_keeper/config"
	"note_keeper/database"
	"os"
	"strconv"
)

var (
	user database.User
	err  error
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	log.Println("------------------- Server start -------------------")
	setDbHost()

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		user_login := r.URL.Query().Get("login")
		user_pass := r.URL.Query().Get("password")

		err = user.Get(user_login, GetHashFromPassword(user_pass))
		if err != nil {
			NotFoundHandler(w, "User not found")
		} else {
			fmt.Fprintf(w, "Hello, %s!\nYou have been authorized.", user.Name)
		}
	})

	http.HandleFunc("/note", func(w http.ResponseWriter, r *http.Request) {
		if !user.IsEmpty() {
			switch r.Method {
			case "GET":
				user.SortNotes("desc")
				userJson, err := json.Marshal(user.Notes)
				CheckError(err)
				JsonResponse(w, userJson)
			case "POST":
				noteText := r.URL.Query().Get("note_text")
				newNote := user.AddNote(noteText)
				noteJson, err := json.Marshal(newNote)
				CheckError(err)
				JsonResponse(w, noteJson)
			case "DELETE":
				noteId, err := strconv.Atoi(r.URL.Query().Get("note_id"))
				CheckError(err)
				err = user.DeleteNoteById(noteId)
				CheckError(err)
				deleteJson, err := json.Marshal(struct {
					DeletedNoteId int `json:"delete_note_id"`
				}{
					DeletedNoteId: noteId,
				})
				CheckError(err)
				JsonResponse(w, deleteJson)
			}
		} else {
			NotFoundHandler(w, "Authorization is required!")
		}
	})

	http.HandleFunc("/note/first_and_last", func(w http.ResponseWriter, r *http.Request) {
		if !user.IsEmpty() {
			switch r.Method {
			case "GET":

				userJson, err := json.Marshal(struct {
					Old_note  database.Note `json:"old_note"`
					Last_note database.Note `json:"last_note"`
				}{
					Old_note:  user.GetOldNotes(1)[0],
					Last_note: user.GetLastNotes(1)[0],
				})
				CheckError(err)
				JsonResponse(w, userJson)
			}
		} else {
			NotFoundHandler(w, "Authorization is required!")
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func setDbHost() {
	database_ip, ok := os.LookupEnv("DB_HOST")
	if ok {
		config.DB_conf.Host = database_ip
		log.Printf("New database host has been set. New host = %s", database_ip)
	} else {
		log.Printf("Database host = %s", config.DB_conf.Host)
	}
}
