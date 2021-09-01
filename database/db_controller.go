package database

import "fmt"

func GetUser(login, password string) (User, error) {
	var user User

	q := fmt.Sprintf(`SELECT id, name, login, expired_at FROM %s WHERE login = '%s' and password = '%s';`,
		user_tabel, login, password)
	err := GetConn().QueryRow(q).Scan(&user.ID, &user.Name, &user.Login, &user.Expired_at)

	return user, err
}

func GetNotesByUser(userId int) ([]Note, error) {
	var notes []Note
	q := fmt.Sprintf(`SELECT id, text_note FROM %s WHERE user_id = %d;`,
		notes_tabel, userId)

	//fmt.Printf("User_%d  Notes:\n", userId)
	rows, err := GetConn().Query(q)
	defer rows.Close()

	if err == nil {
		for rows.Next() {
			var temp_n Note
			if err := rows.Scan(&temp_n.ID, &temp_n.Text); err != nil {
				CheckError(err)
			}
			notes = append(notes, temp_n)
			//fmt.Printf("\tNote_%d: %s\n", temp_n.ID, temp_n.Text)
		}
	}

	return notes, err
}
