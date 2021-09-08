package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"note_keeper/models"
	"strconv"
)

var (
	user models.User
	err  error
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	log.Println("------------------- Server start -------------------")

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
				newNote, err := user.AddNote(noteText)
				if err == nil {
					noteJson, err := json.Marshal(newNote)
					CheckError(err)
					JsonResponse(w, noteJson)
				} else {
					InnerErrorHandler(w, err)
				}
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

	http.HandleFunc("/note/first", func(w http.ResponseWriter, r *http.Request) {
		if !user.IsEmpty() {
			switch r.Method {
			case "GET":

				userJson, err := json.Marshal(struct {
					Old_note models.Note `json:"old_note"`
				}{
					Old_note: user.GetOldNotes(1)[0],
				})
				CheckError(err)
				JsonResponse(w, userJson)
			}
		} else {
			NotFoundHandler(w, "Authorization is required!")
		}
	})

	http.HandleFunc("/note/last", func(w http.ResponseWriter, r *http.Request) {
		if !user.IsEmpty() {
			switch r.Method {
			case "GET":

				userJson, err := json.Marshal(struct {
					Last_note models.Note `json:"last_note"`
				}{
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
