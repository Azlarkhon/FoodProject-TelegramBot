package main

import (
	"telegram-bot/bot"
	"telegram-bot/database"

	"github.com/gin-gonic/gin"
)

func main() {

	database.DBinstance()
	defer database.CloseDB()

	go bot.StartBot()

	router := gin.Default()

	router.Run(":8080")
}
