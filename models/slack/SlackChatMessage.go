package slack

// PostChatMessage представляет сообщение, которое можно запостить в чат слака.
type PostChatMessage struct {
	Text       string `json:"text"`
	Channel    string `json:"channel"`
	IsMarkdown bool   `json:"mrkdwn"`
}
