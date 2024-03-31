package constvar

const (
	APP_ABOUT string = "A serverless Telegram bot"
	APP_SRC   string = "https://github.com/AnnonaOrg/annona_bot"
)

func Version() string {
	return APP_VERSION
}

func Usage() string {
	return APP_ABOUT + "\n" + "I am Annona! TG频道: @annonaOrg"
}
func About() string {
	text := "Annono Bot v" + APP_VERSION + "\n" + Usage()
	return text
}
