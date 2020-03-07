package telegramBot

import (
	"fmt"
	"log"
	"mai_scheduler/cmd"
	"reflect"
	"strings"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func msgError(chatID int64, err error) tgbotapi.MessageConfig {
	return tgbotapi.NewMessage(chatID, err.Error())
}

func TelegramBot(c *cmd.MaiClient, token string) {

	fmt.Println("token =", token)
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for upd := range updates {
		if upd.Message == nil {
			continue
		}

		if reflect.TypeOf(upd.Message.Text).Kind() == reflect.String && upd.Message.Text != "" {
			switch upd.Message.Text {
			case "/start":
				{

					//Отправлем сообщение
					msg := tgbotapi.NewMessage(upd.Message.Chat.ID, "Hi, i'm a wikipedia bot, i can search information in a wikipedia, send me something what you want find in Wikipedia.")
					bot.Send(msg)
				}
			case "/groups":
				{

					info, err := c.GetGroups(8, 3)
					if err != nil {
						bot.Send(msgError(upd.Message.Chat.ID, err))
					}

					msg := tgbotapi.NewMessage(upd.Message.Chat.ID, fmt.Sprintf("найдены группы: ", strings.Join(info.Groups, ", ")))

					bot.Send(msg)

				}
			case "/schedule":
				{
					schs, err := c.GetScheduleByGroup("М8О-311Б-17")
					if err != nil {
						bot.Send(msgError(upd.Message.Chat.ID, err))
					}
					var res string
					for _, sch := range schs {
						res = "\n" + sch.Day + " " + sch.ScDay + ":"

						msg := tgbotapi.NewMessage(upd.Message.Chat.ID, res)
						bot.Send(msg)
						for _, row := range sch.Rows {
							res = "\n" + row.StartAt + " - " + row.EndAt
							res += " " + row.Title + " " + row.LessonType + " " + row.Lecturer + " " + row.Location

							msg = tgbotapi.NewMessage(upd.Message.Chat.ID, res)
							bot.Send(msg)
						}
					}
					msg := tgbotapi.NewMessage(upd.Message.Chat.ID, res)

					bot.Send(msg)
				}
			case "/info":
				{
					res := `Привет, студент (8 фака 3 курса и 311 группы)!
					Здесь ты можешь выбрать /groups - чтобы узнать какие группы есть на 8 факе 3 курса и
					загрузить расписание командой /schedule
					`
					msg := tgbotapi.NewMessage(upd.Message.Chat.ID, res)

					bot.Send(msg)
				}
			default:
				{
					res := `Привет, студент (8 фака 3 курса и 311 группы)!
					Здесь ты можешь выбрать /groups - чтобы узнать какие группы есть на 8 факе 3 курса и
					загрузить расписание командой /schedule
					`
					msg := tgbotapi.NewMessage(upd.Message.Chat.ID, res)

					bot.Send(msg)
				}
			}

		}

	}

}
