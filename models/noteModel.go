package models

import (
	"math/rand"
	"time"
)

var (
	randomN = 1000
	random  = rand.New(rand.NewSource(time.Now().UnixNano()))
)

type Note struct {
	ID         int       `json:"id"`
	Created_at time.Time `json:"created_at"`
	Text       string    `json:"text"`
}

func (n *Note) Init(noteText string) {
	n.setId()
	n.Created_at, n.Text = time.Now(), noteText
}

func (n *Note) setId() {
	n.ID = random.Intn(randomN)
}
