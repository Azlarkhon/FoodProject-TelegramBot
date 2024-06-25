package controllers

import (
	"log"
	"telegram-bot/database"
)

func InsertUser(id int) {
	db := database.DBinstance()
	
	var tempID int

	err := db.QueryRow("SELECT id FROM users WHERE id = ?", id).Scan(&tempID)
	if err != nil {
		log.Print("error", err)
	}

	if tempID == 0 {
		_, err := db.Exec("INSERT INTO users (id) VALUES (?)", id)
		if err != nil {
			log.Print("insert error", err)
		}
	}
}