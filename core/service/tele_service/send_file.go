package tele_service

import (
	"bytes"
	"fmt"
	"time"

	"github.com/AnnonaOrg/annona_bot/core/log"

	tele "gopkg.in/telebot.v3"
)

func SendFileWithString(c tele.Context, fileBodyStr string, caption string) {

	fileBody := []byte(fileBodyStr)
	fileMsg := &tele.Document{
		File: tele.FromReader(
			bytes.NewReader(fileBody),
		),
		FileName: fmt.Sprintf("all_%s.txt", time.Now().Format(time.DateOnly)),
	}
	fileMsg.Caption = caption

	if err := c.Reply(fileMsg, tele.ModeHTML, tele.NoPreview); err != nil {
		log.Errorf("Reply(%d): %v", c.Sender().ID, err)
	}

}
