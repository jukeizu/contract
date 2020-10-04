package contract

type Reaction struct {
	ChannelId string `json:"channelId,omitempty" yaml:"channelId,omitempty"`
	MessageId string `json:"messageId,omitempty" yaml:"messageId,omitempty"`
	EmojiId   string `json:"emojiId,omitempty" yaml:"emojiId,omitempty"`
}
