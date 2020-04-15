package entity

import "context"

type channelStruct struct {
	*model
	accountId *int64
	name      *string
}

func newChannel() *channel {
	return &channel{channelStruct: newChannelStruct()}
}

func newChannelStruct() *channelStruct {
	return &channelStruct{model: newModel()}
}

type channelsStruct []channelInterface

type channelInterface interface {
	Model
	AccountId() *int64
	Name() *string
	SetAccountId(accountId *int64)
	SetName(name *string)
	BeforeRead(ctx context.Context) error
	BeforeAdd(ctx context.Context) error
}
type channelsInterface interface {
	Len() int
	Append(req channelInterface)
	Elements() []channelInterface
}

func (m *channelsStruct) Len() int {
	if m != nil {
		return len(*m)
	}
	return 0
}
func (m *channelsStruct) Append(req channelInterface) {
	if m != nil {
		*m = append(*m, req)
	}
}
func (m *channelsStruct) Elements() []channelInterface {
	return *m
}
func (m *channelStruct) AccountId() *int64 {
	return m.accountId
}

func (m *channelStruct) Name() *string {
	return m.name
}

func (m *channelStruct) SetAccountId(accountId *int64) {
	m.accountId = accountId
}

func (m *channelStruct) SetName(name *string) {
	m.name = name
}

func (m *channelStruct) BeforeRead(ctx context.Context) error {
	return nil
}

func (m *channelStruct) BeforeAdd(ctx context.Context) error {
	return nil
}
