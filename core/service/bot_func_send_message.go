package service

import (
	log "github.com/sirupsen/logrus"
	tele "gopkg.in/telebot.v3"
)

func SendMessage(botToken string, reciverId int64, m interface{}, parseMode tele.ParseMode, noButton bool, button interface{}) error {
	bot, err := tele.NewBot(tele.Settings{
		Token:       botToken,
		Synchronous: false,
	})
	// common.Must(err)
	if err != nil {
		return err
	}

	reciver := &tele.User{
		ID: reciverId,
	}
	if noButton {
		if _, err := bot.Send(reciver, m, parseMode, tele.NoPreview); err != nil {
			log.Errorf("Send(%s,%d,%#v,%v) Msg Error: %v", botToken, reciverId, m, parseMode, err)
			return err
		}
	} else {
		if _, err := bot.Send(reciver, m, parseMode, button, tele.NoPreview); err != nil {
			log.Errorf("Send(%s,%d,%#v,%v,%+v) Msg Error: %v", botToken, reciverId, m, parseMode, button, err)
			return err
		}
	}

	return nil
}
