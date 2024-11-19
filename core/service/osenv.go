package service

import (
	"strconv"
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

func IsEnableFilterSameSenderUserMsg() bool {
	// Filter messages triggered multiple times by the same user within a short period of time
	return strings.EqualFold(osenv.Getenv("FILTER_SAME_SENDER_USER_MSG_ENABLE"), "true")
}
func GetMinCountFilterSameSenderUserMsg() int {
	item := osenv.Getenv("FILTER_SAME_SENDER_USER_MSG_MINI_COUNT")
	count, err := strconv.Atoi(item)
	if err != nil || count == 0 {
		return 500
	}
	return count
}
