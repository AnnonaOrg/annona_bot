package service

import (
	"strings"

	"github.com/AnnonaOrg/osenv"
)

func GetURLCardbot() string {
	return osenv.Getenv("BUTTON_LINK_Cardbot")
}
func GetURLSubmitNewGroup() string {
	return osenv.Getenv("BUTTON_LINK_SubmitNewGroup")
}

// _retry
func IsRetryPushMsgEnable() bool {
	return strings.EqualFold(osenv.Getenv("RETRY_PUSH_MSG_ENABLE"), "true")
}
