package entity

type channel struct {
	*channelStruct
}

type channels struct {
	channelsStruct
}

type Channel interface {
	channelInterface
}
type Channels interface {
	channelsInterface
}

func NewChannel() Channel {
	return newChannel()
}

func NewChannels() Channels {
	return &channels{}
}
