package database

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Login string `json:"login"`
	//Password string		`json:"password"`
	Registered_at time.Time `json:"registered_at"`
	Expired_at    time.Time `json:"expired_at"`
	Notes         []Note    `json:"notes"`
}

func (u *User) IsEmpty() bool {
	return u.ID == 0 && u.Login == "" && u.Registered_at.IsZero()
}

func (u *User) Get(login, password string) error {
	q := fmt.Sprintf(`SELECT id, name, login, expired_at FROM %s WHERE login = '%s' and password = '%s';`,
		user_tabel, login, password)
	err := GetConn().QueryRow(q).Scan(&u.ID, &u.Name, &u.Login, &u.Expired_at)

	if err == nil {
		u.Notes, err = getNotesByUserId(u.ID)
	}

	return err
}

func (u *User) PrintNotes() {
	fmt.Printf("User_%d  %s  Notes:\n", u.ID, u.Name)
	for _, note := range u.Notes {
		fmt.Printf("\tNote_%d [%s]:  %s\n", note.ID, note.Created_at.Format("2006-01-02 15:04:05"), note.Text)
	}
}

func (u *User) AddNote(textNote string) *Note {
	var newNote Note
	err := newNote.Init(u.ID, textNote)
	if err != nil {
		panic(err)
	}

	u.Notes = append(u.Notes, newNote)

	return &newNote
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
		err := note.delete(u.ID)
		if err == nil {
			u.Notes = append(u.Notes[:id_in_slice], u.Notes[id_in_slice+1:]...)
		}
		return err
	}

	return nil
}

func copySlice(arr []Note) []Note {
	ans := make([]Note, len(arr))
	copy(ans, arr)
	return ans
}
