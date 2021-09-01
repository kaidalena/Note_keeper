package api

import (
	"note_keeper/database"
)

func Login(login, password string) (database.User, error) {
	user, err := database.GetUser(login, password)

	if err == nil {
		user.Notes, err = database.GetNotesByUser(user.ID)
	}

	return user, err
}
