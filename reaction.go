package contract

type Reaction struct {
	ChannelId string `json:"channelId,omitempty"`
	MessageId string `json:"messageId,omitempty"`
	EmojiId   string `json:"emojiId,omitempty"`
}
