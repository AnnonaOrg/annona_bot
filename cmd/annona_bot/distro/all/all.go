package all

import (
	_ "github.com/AnnonaOrg/annona_bot/core/log"

	_ "github.com/AnnonaOrg/annona_bot/core/features/about"
	_ "github.com/AnnonaOrg/annona_bot/core/features/ping"
	_ "github.com/AnnonaOrg/annona_bot/core/features/start"

	_ "github.com/AnnonaOrg/annona_bot/core/features/callback"

	_ "github.com/AnnonaOrg/annona_bot/core/features/card_features"
	_ "github.com/AnnonaOrg/annona_bot/core/features/keyword_features"
	_ "github.com/AnnonaOrg/annona_bot/core/features/telebot_features"
	_ "github.com/AnnonaOrg/annona_bot/core/features/user_features"

	_ "github.com/AnnonaOrg/annona_bot/core/features/blockformchatid_features"
	_ "github.com/AnnonaOrg/annona_bot/core/features/blockformsenderid_features"
	_ "github.com/AnnonaOrg/annona_bot/core/features/blockword_features"

	_ "github.com/AnnonaOrg/annona_bot/core/features/text"

	_ "github.com/AnnonaOrg/annona_bot/core/features/newbot_features"

	_ "github.com/AnnonaOrg/annona_bot/core/features/keyword_history_features"
)
