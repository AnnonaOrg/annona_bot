package service

import (
	"github.com/AnnonaOrg/osenv"
)

func GetURLCardbot() string {
	return osenv.Getenv("BUTTON_LINK_Cardbot")
}
func GetURLSubmitNewGroup() string {
	return osenv.Getenv("BUTTON_LINK_SubmitNewGroup")
}
