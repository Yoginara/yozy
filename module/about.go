package module

import (
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"hanacore/utils/console"
	"strings"
)

type AboutModule struct{}

func (m *AboutModule) Handle(ctx context.Context, b *bot.Bot, update *models.Update) {
	moduleName := "About"
	moduleCommand := "/about"
	senderID := bot.EscapeMarkdown(fmt.Sprintf("%d", update.Message.From.ID))

	message := update.Message.Text
	if strings.HasPrefix(message, moduleCommand) {
		kb := &models.InlineKeyboardMarkup{
			InlineKeyboard: [][]models.InlineKeyboardButton{
				{
					{Text: "GitHub - herobuxx/hana", URL: "https://github.com/herobuxx/hana"},
				},
			},
		}

		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:      update.Message.Chat.ID,
			Text:        "Yozy is a Telegram Bot designed to be modular and easy to develop. Check out my source code here!",
			ReplyMarkup: kb,
		})
		console.ShowLog(moduleName, senderID)
	}
}

func init() {
	RegisterModule(&AboutModule{})
}
