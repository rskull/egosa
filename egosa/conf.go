package egosa

type Config struct {
	Core     SectionCore     `toml:core`
	Twitter  SectionTwitter  `toml:"twitter"`
	Chatwork SectionChatwork `toml:"chatwork"`
	Slack    SectionSlack    `toml:"slack"`
}

type SectionCore struct {
	IntervalSec int64 `toml:"intervalSec"`
	Count       int   `toml:"count"`
}

type SectionTwitter struct {
	ConsumerKey       string `toml:"consumerKey"`
	ConsumerKeySecret string `toml:"consumerKeySecret"`
	AuthKey           string `toml:"authKey"`
	AuthKeySecret     string `toml:"authKeySecret"`
	SearchQuery       string `toml:"searchQuery"`
	ResultType        string `toml:"resultType"`
	Lang              string `toml:"lang"`
}

type SectionChatwork struct {
	ApiKey   string `toml:"apiKey"`
	RoomID   string `toml:"roomID"`
	SendBody string `toml:"sendBody"`
	Enable   bool   `toml:"enable"`
}

type SectionSlack struct {
	ApiKey  string `toml:"apiKey"`
	Channel string `toml:"channel"`
	Enable  bool   `toml:"enable"`
}
