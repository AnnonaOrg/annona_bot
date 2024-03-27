package all

import (
	_ "github.com/AnnonaOrg/annona_bot/internal/log"

	_ "github.com/AnnonaOrg/annona_bot/features/about"
	_ "github.com/AnnonaOrg/annona_bot/features/ping"
	_ "github.com/AnnonaOrg/annona_bot/features/start"

	_ "github.com/AnnonaOrg/annona_bot/features/callback"

	_ "github.com/AnnonaOrg/annona_bot/features/card_features"
	_ "github.com/AnnonaOrg/annona_bot/features/keyword_features"
	_ "github.com/AnnonaOrg/annona_bot/features/telebot_features"
	_ "github.com/AnnonaOrg/annona_bot/features/user_features"

	_ "github.com/AnnonaOrg/annona_bot/features/blockformchatid_features"
	_ "github.com/AnnonaOrg/annona_bot/features/blockformsenderid_features"
	_ "github.com/AnnonaOrg/annona_bot/features/blockword_features"

	_ "github.com/AnnonaOrg/annona_bot/features/text"

	_ "github.com/AnnonaOrg/annona_bot/features/newbot_features"
)
