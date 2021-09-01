package database

import (
	"fmt"
	"time"
)

type Note struct {
	ID int `json:"id"`
	//UserId int			`json:"user_id"`
	Text       string    `json:"note_text"`
	Created_at time.Time `json:"created_at"`
}

func (n *Note) Init(userId int, noteText string) error {
	q := fmt.Sprintf(`INSERT INTO %s (user_id, text_note) values (%d, '%s') RETURNING id, created_at;`,
		notes_tabel, userId, noteText)
	err := GetConn().QueryRow(q).Scan(&n.ID, &n.Created_at)

	if err == nil {
		n.Text = noteText
	}

	return err
}

func (n *Note) delete(userId int) error {
	q := fmt.Sprintf(`DELETE FROM %s WHERE user_id = %d and id = %d;`,
		notes_tabel, userId, n.ID)
	_, err := GetConn().Exec(q)

	return err
}

func getNotesByUserId(userId int) ([]Note, error) {
	var notes []Note
	q := fmt.Sprintf(`SELECT id, text_note, created_at FROM %s WHERE user_id = %d;`,
		notes_tabel, userId)

	//fmt.Printf("User_%d  Notes:\n", userId)
	rows, err := GetConn().Query(q)
	defer rows.Close()

	if err == nil {
		for rows.Next() {
			var temp_n Note
			if err := rows.Scan(&temp_n.ID, &temp_n.Text, &temp_n.Created_at); err != nil {
				CheckError(err)
			}
			notes = append(notes, temp_n)
			//fmt.Printf("\tNote_%d: %s\n", temp_n.ID, temp_n.Text)
		}
	}

	return notes, err
}
