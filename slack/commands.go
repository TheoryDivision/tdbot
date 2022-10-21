package slack

var CommandChan = make(chan CommandInfo)

type CommandInfo struct {
	Fields    []string `json:"fields"`
	Channel   string   `json:"channel"`
	TimeStamp string   `json:"timestamp"`
}

func (sc *slackClient) RunSocketMode() {
	sc.client.Run()
}

func (sc *slackClient) getChannelName(channelID string) (channel string) {
	ch, err := sc.api.GetConversationInfo(channelID, false)
	if err != nil {
		sc.logger.WithField("err", err).Error("Couldn't find conversation info.")
	}
	return ch.Name
}

func (sc *slackClient) getUserName(userID string) (user string) {
	us, err := sc.api.GetUserInfo(userID)
	if err != nil {
		sc.logger.WithField("err", err).Error("Couldn't find conversation info.")
	}
	return us.Profile.DisplayName
}
