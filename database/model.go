package database

import (
	"time"
)

type Note struct {
	ID   int
	Text string
}

type User struct {
	ID    int
	Name  string
	Login string
	//Password string
	Expired_at time.Time // user.Expired_at.Format("2006-01-02 15:04:05")
	Notes      []Note
}
