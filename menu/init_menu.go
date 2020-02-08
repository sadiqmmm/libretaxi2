package menu

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"libretaxi/context"
	"libretaxi/objects"
	"log"
)

type InitMenuHandler struct {
}

func (handler *InitMenuHandler) Handle(user *objects.User, context *context.Context, message *tgbotapi.Message) {
	log.Println("Init menu")

	// Send welcome message
	msg := tgbotapi.NewMessage(user.UserId, user.Locale().Get("init_menu.welcome"))
	context.Send(msg)

	user.MenuId = objects.Menu_AskLocation
	context.Repo.SaveUser(user)
}

func NewInitMenu() *InitMenuHandler {
	handler := &InitMenuHandler{}
	return handler
}
