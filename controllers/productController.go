package controllers

import (
	"log"
	"telegram-bot/database"
	"telegram-bot/models"
)

func GetFood(name string) models.Food {
	db := database.DBinstance()

	var food models.Food

	err := db.QueryRow("SELECT name, description, price FROM foods WHERE name = ?", name).Scan(
		&food.Name, &food.Description, &food.Price)
	if err != nil {
		log.Print("Error: ", err)
	}

	log.Print(food)

	return food
}

func FetchFoodNames() []string {
	db := database.DBinstance()

	rows, err := db.Query("SELECT name FROM foods")
	if err != nil {
		return nil
	}
	defer rows.Close()

	var foodNames []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil
		}
		foodNames = append(foodNames, name)
	}
	return foodNames
}
