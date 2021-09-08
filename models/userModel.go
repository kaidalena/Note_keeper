package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
	"time"
)

var (
	fileUsersPath = "data.json"
)

type User struct {
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	Login         string    `json:"login"`
	Password      string    `json:"password"`
	Registered_at time.Time `json:"registered_at"`
	Expired_at    time.Time `json:"expired_at"`
	Notes         []Note    `json:"notes"`
}

func (u *User) IsEmpty() bool {
	return u.ID == 0 && u.Login == "" && u.Registered_at.IsZero()
}

func (u *User) Get(login, password string) error {
	users, err := getUsersFromFile(fileUsersPath)
	if err != nil {
		return err
	}

	for _, user := range users {
		if user.Login == login && user.Password == password {
			*u = user
			break
		}
	}

	for i := range u.Notes {
		u.Notes[i].setId()
		fmt.Printf("%v set id %v\n", u.Notes[i].Text, u.Notes[i].ID)
	}

	return err
}

func (u *User) PrintNotes() {
	fmt.Printf("User_%d  %s  Notes:\n", u.ID, u.Name)
	for _, note := range u.Notes {
		fmt.Printf("\tNote_%d [%s]:  %s\n", note.ID, note.Created_at.Format("2006-01-02 15:04:05"), note.Text)
	}
}

func (u *User) AddNote(textNote string) (*Note, error) {
	var newNote Note
	newNote.Init(textNote)

	u.Notes = append(u.Notes, newNote)

	return &newNote, u.updateUserInFile(fileUsersPath)
}

func (u *User) GetNoteById(noteId int) (*Note, int) {
	for i, note := range u.Notes {
		if note.ID == noteId {
			return &note, i
		}
	}

	return nil, 0
}

func (u *User) SortNotes(direction string) {
	switch strings.ToUpper(direction) {
	case "ASC":
		sort.SliceStable(u.Notes, func(i, j int) bool {
			return u.Notes[i].Created_at.Before(u.Notes[j].Created_at)
		})
	case "DESC":
		sort.SliceStable(u.Notes, func(i, j int) bool {
			return u.Notes[i].Created_at.After(u.Notes[j].Created_at)
		})
	}
}

func (u *User) GetOldNotes(count int) []Note {
	u.SortNotes("Asc")
	return copySlice(u.Notes[:count])
}

func (u *User) GetLastNotes(count int) []Note {
	u.SortNotes("desc")
	return copySlice(u.Notes[:count])
}

func (u *User) DeleteNoteById(noteId int) error {
	note, id_in_slice := u.GetNoteById(noteId)
	if note != nil {
		u.Notes = append(u.Notes[:id_in_slice], u.Notes[id_in_slice+1:]...)
	}

	return u.updateUserInFile(fileUsersPath)
}

func copySlice(arr []Note) []Note {
	ans := make([]Note, len(arr))
	copy(ans, arr)
	return ans
}

func getUsersFromFile(filePath string) ([]User, error) {
	jsonFile, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()

	jsonData, err := ioutil.ReadAll(jsonFile)
	var users []User
	err = json.Unmarshal(jsonData, &users)

	return users, err
}

func saveUsersInFile(filePath string, users []User) error {
	jsonFile, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()

	usersJson, err := json.Marshal(users)
	err = ioutil.WriteFile(filePath, usersJson, 0644)

	return err
}

func (u *User) updateUserInFile(filePath string) error {
	users, err := getUsersFromFile(filePath)
	if err != nil {
		return err
	}

	index := -1
	for i, user := range users {
		if user.ID == u.ID {
			index = i
			break
		}
	}

	if index >= 0 {
		users[index] = *u
	}

	return saveUsersInFile(filePath, users)
}
