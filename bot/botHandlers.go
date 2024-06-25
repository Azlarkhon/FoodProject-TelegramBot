package bot

import (
	"log"
	"strconv"
	controller "telegram-bot/controllers"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var foodImageURLs = map[string]string{
	"Hot-dog": "images/Hot-dog.jpeg",
	"Lavash":  "images/Lavash.jpg",
	"Burger":  "images/Burger.jpg",
	"Shourma": "images/Shourma.jpg",
	"Pizza":  "images/Pizza.jpg",
	"Sushi": "images/Sushi.jpg",
	"Chicken": "images/Chicken.jpeg",
	"Roast beef" : "images/Roast-beef.jpg",
	"Sandwich" : "images/Sandwich.jpg",
}

var tempFood string

func handleMessage(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	var msg tgbotapi.MessageConfig

	userID := message.From.ID
	username := message.From.FirstName

	chatID := message.Chat.ID

	switch message.Text {
	case "/start", "Back to main page":
		controller.InsertUser(userID)

		msg = tgbotapi.NewMessage(chatID, "Main page")
		replyKeyboard := tgbotapi.NewReplyKeyboard(
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton("Menu"),
			),
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton("My orders"),
			),
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton("My info"),
			),
		)
		msg.ReplyMarkup = replyKeyboard

	case "Menu", "Back to menu":
		foodNames := controller.FetchFoodNames()
		keyboard := CreateKeyboard(foodNames)

		msg = tgbotapi.NewMessage(chatID, "Please select the food you want.")
		msg.ReplyMarkup = keyboard

		replyKeyboard := tgbotapi.NewReplyKeyboard(
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton("My cart"),
			),
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton("Back to main page"),
			),
		)

		// Append the additional reply keyboard to the existing one
		msg.ReplyMarkup = tgbotapi.ReplyKeyboardMarkup{
			Keyboard:       append(msg.ReplyMarkup.(tgbotapi.ReplyKeyboardMarkup).Keyboard, replyKeyboard.Keyboard...),
			ResizeKeyboard: true,
		}

	case "My orders":
		msg = tgbotapi.NewMessage(chatID, "Here are your orders: [List of orders]")
	case "My cart":
		msg = tgbotapi.NewMessage(chatID, "Your cart is currently empty.")
	case "Order":

	case "My info":
		msg = tgbotapi.NewMessage(chatID, "Your id: "+strconv.Itoa(int(userID))+"\nYour name is: "+username)
	default:
		imageURL, ok := foodImageURLs[message.Text]
		if ok {
			sendFoodDetail(bot, chatID, message.Text, imageURL)
			return
		}
		msg = tgbotapi.NewMessage(chatID, "I don't understand that command.")
	}

	if msg.Text != "" {
		if _, err := bot.Send(msg); err != nil {
			log.Printf("Error sending message: %v", err)
		}
	}
}

func handleCallbackQuery(bot *tgbotapi.BotAPI, callbackQuery *tgbotapi.CallbackQuery) {
	var responseText string
	switch callbackQuery.Data {
	case "decrease_" + tempFood:
		responseText = "You decreased " + tempFood + " quantity."
	case "increase_" + tempFood:
		responseText = "You increased " + tempFood + " quantity."
	default:
		responseText = "Unknown action."
	}

	msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, responseText)
	if _, err := bot.Send(msg); err != nil {
		log.Printf("Error sending message: %v", err)
	}

	// Optionally, you can send an answer callback query to remove the loading animation on the button
	answerCallback := tgbotapi.NewCallback(callbackQuery.ID, "")
	if _, err := bot.AnswerCallbackQuery(answerCallback); err != nil {
		log.Printf("Error answering callback query: %v", err)
	}
}

func sendFoodDetail(bot *tgbotapi.BotAPI, chatID int64, foodName string, imageURL string) {
	tempFood = ""
	tempFood = foodName
	photo := tgbotapi.NewPhotoUpload(chatID, imageURL)

	food := controller.GetFood(foodName)

	photo.Caption = foodName + "\n" + food.Description + "\nPrice: " + strconv.FormatFloat(food.Price, 'f', 2, 64)
	photo.ReplyMarkup = getFoodDetailInlineKeyboard(foodName)
	if _, err := bot.Send(photo); err != nil {
		log.Printf("Error sending message: %v", err)
	}
}

func getFoodDetailInlineKeyboard(food string) tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("- ", "decrease_"+food),
			tgbotapi.NewInlineKeyboardButtonData("+ ", "increase_"+food),
		),
	)
}

func CreateKeyboard(foodNames []string) tgbotapi.ReplyKeyboardMarkup {
	var rows [][]tgbotapi.KeyboardButton

	log.Print(foodNames)

	for i := 0; i < len(foodNames); i += 3 {
		var row []tgbotapi.KeyboardButton
		for j := i; j < i+3 && j < len(foodNames); j++ {
			btn := tgbotapi.NewKeyboardButton(foodNames[j])
			row = append(row, btn)
		}
		rows = append(rows, tgbotapi.NewKeyboardButtonRow(row...))
	}

	keyboard := tgbotapi.NewReplyKeyboard(rows...)
	return keyboard
}
