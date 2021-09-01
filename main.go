package main

import (
	"fmt"
	"log"
	"note_keeper/api"
	"note_keeper/config"
	"os"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	log.Println("------------------- Server start -------------------")

	setDbHost()
	user, err := api.Login("little_coon", api.GetHashFromPassword("pwd"))
	//fmt.Println(user)
	CheckError(err)

	user.PrintNotes()
	newNoteId := user.AddNote("Note from app")
	fmt.Printf("Id new note = %d\n", newNoteId)

	user.PrintNotes()
	CheckError(user.DeleteNoteById(newNoteId))
	user.PrintNotes()
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
