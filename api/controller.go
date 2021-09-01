package api

import (
	"note_keeper/database"
)

func Login(login, password string) (database.User, error) {
	var user database.User

	err := user.Get(login, password)

	return user, err
}

func SortNotes(user database.User, direction string) {
	user.SortNotes(direction)
}
