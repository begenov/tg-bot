package handlers

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/begenov/tg-bot/internal/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (api *TelegramAPI) profileUser(update tgbotapi.Update, chatId int64, msg tgbotapi.MessageConfig) {
	switch api.usermapa[chatId].Stage {
	case 0:
		if update.CallbackQuery != nil {
			api.choseKazakhHandler(update, msg, chatId)
			break
		}
	case 1:
		if update.Message != nil {
			api.nameHandler(update, chatId, msg)
			break
		}
	case 2:
		if update.Message.Contact != nil {
			api.phoneNumberHandler1(update, chatId, msg)
			break
		}
		if update.Message != nil {
			api.phoneNumberHandler2(update, chatId, msg)
			break
		}
	case 4:
		if update.Message != nil {
			api.checkPhoneNumberHandler(update, chatId, msg)
			break

		}
		// fallthrough
	case 5:
		if update.CallbackQuery != nil {
			api.coverLetterHandler(update, chatId, msg)
			break
		}
	case 6:
		if update.Message != nil {
			msg := tgbotapi.NewMessage(chatId, update.Message.Text)
			api.ageUserHandler(update, chatId, msg)
			break
		}

	case 7:

		if update.CallbackQuery != nil {
			api.genderHandler(update, chatId, msg)
			break
		}
	case 8:
	default:
		// error
	}
}

func (api *TelegramAPI) choseKazakhHandler(update tgbotapi.Update, msg tgbotapi.MessageConfig, chatId int64) {
	lang := update.CallbackQuery.Data
	api.usermapa[chatId].Stage = 1
	api.usermapa[chatId].Lang = lang

	msg2 := tgbotapi.NewMessage(chatId, "")
	if lang == models.Kazakh {
		msg.Text = models.ChoseKazakh
		msg2.Text = models.KazakhName
	} else if lang == models.Russian {
		msg.Text = models.ChoseRussian
		msg2.Text = models.RussianName
	}

	api.bot.Send(msg)
	api.bot.Send(msg2)
}

// share telefon

func (api *TelegramAPI) phoneNumberHandler1(update tgbotapi.Update, chatId int64, msg tgbotapi.MessageConfig) {
	phoneNumber := update.Message.Contact.PhoneNumber
	// log.Fatal(phoneNumber)
	api.usermapa[chatId].Phone = phoneNumber
	api.usermapa[chatId].Stage = 4
	if api.usermapa[chatId].Lang == models.Kazakh {
		msg.Text = models.KazakhNumberInfo
	} else {
		msg.Text = models.RussianNumberInfo
	}
	api.bot.Send(msg)
}

func (api *TelegramAPI) phoneNumberHandler2(update tgbotapi.Update, chatId int64, msg tgbotapi.MessageConfig) {
	phoneNumber := update.Message.Text
	api.usermapa[chatId].Phone = phoneNumber
	api.usermapa[chatId].Stage = 4
	if api.usermapa[chatId].Lang == models.Kazakh {
		msg.Text = models.KazakhNumberInfo
	} else if api.usermapa[chatId].Lang == models.Russian {
		msg.Text = models.RussianNumberInfo
	}
	api.bot.Send(msg)
}

// Share Name

func (api *TelegramAPI) checkPhoneNumberHandler(update tgbotapi.Update, chatId int64, msg tgbotapi.MessageConfig) {
	code := update.Message.Text
	if code == "0000" {
		if api.usermapa[chatId].Lang == models.Kazakh {
			msg.Text = models.KazakhNumberInfo
		} else {
			msg.Text = models.RussianNumberInfo
		}
		// api.usermapa[chatId].Stage = 5
		// // api.bot.Send(msg)
		// continue
	}
	msg1 := tgbotapi.NewMessage(chatId, "")
	work := ""
	employee := ""
	if api.usermapa[chatId].Lang == models.Kazakh {
		msg.Text = models.KazakhNumberRetrieved
		searchText := fmt.Sprintf(models.KazakhSearch, api.usermapa[chatId].Name)
		msg1.Text = searchText
		work = models.KazakhWorkButton
		employee = models.KazakhEmployeeButton
	} else if api.usermapa[chatId].Lang == models.Russian {
		msg.Text = models.RussianNumberRetrieved
		searchText := fmt.Sprintf(models.RussianSearch, api.usermapa[chatId].Name)
		msg1.Text = searchText
		work = models.RussianWorkButton
		employee = models.RussianEmployeeButton
	}

	api.bot.Send(msg)
	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(work, "1"),
			tgbotapi.NewInlineKeyboardButtonData(employee, "0"),
		),
	)
	msg1.ReplyMarkup = inlineKeyboard
	api.usermapa[chatId].Stage = 5
	api.bot.Send(msg1)
}

func (api *TelegramAPI) nameHandler(update tgbotapi.Update, chatId int64, msg tgbotapi.MessageConfig) {
	name := update.Message.Text
	api.usermapa[chatId].Name = name

	shareButton := tgbotapi.NewKeyboardButtonContact(models.RussianNumberButton)
	msg2 := tgbotapi.NewMessage(chatId, "")

	if api.usermapa[chatId].Lang == models.Kazakh {

		msg.Text = models.KazakhHello + name
		msg2.Text = models.KazakhNumberButton
		shareButton.Text = models.KazakhNumberButton
	} else {
		msg.Text = models.RussianHello + name
		msg2.Text = models.RussianNumberButton
		shareButton.Text = models.RussianNumberButton

	}

	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(shareButton),
	)

	api.bot.Send(msg)

	msg2.ReplyMarkup = keyboard
	api.bot.Send(msg2)
	api.usermapa[chatId].Stage = 2
}

func (api *TelegramAPI) coverLetterHandler(update tgbotapi.Update, chatId int64, msg tgbotapi.MessageConfig) {
	api.usermapa[chatId].Aim = update.CallbackQuery.Data

	age := ""
	if api.usermapa[chatId].Lang == models.Kazakh {
		msg.Text = models.KazakhAccompanyingMess
		age = models.KazakhAgeInfo

	} else if api.usermapa[chatId].Lang == models.Russian {
		msg.Text = models.RussianAccompanyingMess
		age = models.RussianAgeInfo
	}
	api.bot.Send(msg)
	time.Sleep(3 * time.Second)
	msg.Text = age
	api.usermapa[chatId].Stage = 6
	api.bot.Send(msg)
}

func (api *TelegramAPI) ageUserHandler(update tgbotapi.Update, chatId int64, msg tgbotapi.MessageConfig) {
	api.usermapa[chatId].Age, _ = strconv.Atoi(msg.Text)

	showGender := ""

	var genderMale, genderFemale string
	if api.usermapa[chatId].Lang == models.Kazakh {
		genderMale = models.KazakhGenderMale
		genderFemale = models.KazakhGenderFemale
		showGender = models.KazakhGender

	} else if api.usermapa[chatId].Lang == models.Russian {
		genderMale = models.RussianGenderMale
		genderFemale = models.RussianGenderFemale
		showGender = models.RussianGender
	}

	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(genderMale, "1"),
			tgbotapi.NewInlineKeyboardButtonData(genderFemale, "0"),
		),
	)

	msg.ReplyMarkup = inlineKeyboard
	msg.Text = showGender
	api.usermapa[chatId].Stage = 7

	api.bot.Send(msg)
}

func (api *TelegramAPI) genderHandler(update tgbotapi.Update, chatId int64, msg tgbotapi.MessageConfig) {
	gen, err := strconv.Atoi(update.CallbackQuery.Data)
	if err != nil {
		fmt.Println("gender handler error --------")
		log.Fatal(err)

	}
	// gender := ""
	info := ""
	if api.usermapa[chatId].Lang == models.Kazakh {

		info = models.InfoInKazakh

	} else if api.usermapa[chatId].Lang == models.Russian {

		info = models.InfoInRussian
	}

	api.usermapa[chatId].Gender = gen
	api.usermapa[chatId].Stage = 8
	ms := fmt.Sprintf(info, api.usermapa[chatId].Name, api.usermapa[chatId].Phone, api.usermapa[chatId].Aim, api.usermapa[chatId].Age, api.usermapa[chatId].Gender)
	if api.usermapa[chatId].Lang == models.Kazakh {
		msg.Text = ms
	} else if api.usermapa[chatId].Lang == models.Russian {
		msg.Text = ms
	}
	go func() {
		api.usermapa[chatId].ChatID = int(chatId)
		err := api.services.User.Create(context.Background(), *api.usermapa[chatId])
		log.Println(err)
	}()
	api.bot.Send(msg)
}

func (api *TelegramAPI) Hello(message *tgbotapi.Message, chatId int64) {
	msg := tgbotapi.NewMessage(chatId, models.InfoTelega)
	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(models.Kazakh, models.Kazakh),
			tgbotapi.NewInlineKeyboardButtonData(models.Russian, models.Russian),
		),
	)
	msg.ReplyMarkup = inlineKeyboard

	api.usermapa[chatId] = &models.User{Stage: 0}

	api.bot.Send(msg)
	if api.usermapa[chatId].Aim == "1" {
		api.usermapa[chatId].Stage = 8
		msg.Text = "В какой сфере вы бы хотели найти работу?"
		inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Торговля", "0"),
				tgbotapi.NewInlineKeyboardButtonData("Общепит", "1"),
				tgbotapi.NewInlineKeyboardButtonData("Другое", "2"),
				tgbotapi.NewInlineKeyboardButtonData("Пропустить шаг", "3"),
			),
		)

		msg.ReplyMarkup = inlineKeyboard
		api.bot.Send(msg)
	}
}

func (api *TelegramAPI) workFieldHandler(update tgbotapi.Update, chatId int64, msg tgbotapi.MessageConfig) {
	workField, err := strconv.Atoi(update.CallbackQuery.Data)
	if err != nil {
		fmt.Println("WorkField handler error --------")
		log.Fatal(err)
	}
	log.Println(workField, "------------------------------------------")
	msg.Text = "Дальше будеь выбор профессиии, пока в разработке"
	api.bot.Send(msg)
}
