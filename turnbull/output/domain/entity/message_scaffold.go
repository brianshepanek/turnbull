package entity

import "context"

type messageStruct struct {
	*model
	accountId *int64
	channelId *int64
	userId    *int64
	message   *string
}

func newMessage() *message {
	return &message{messageStruct: newMessageStruct()}
}

func newMessageStruct() *messageStruct {
	return &messageStruct{model: newModel()}
}

type messagesStruct []messageInterface

type messageInterface interface {
	Model
	AccountId() *int64
	ChannelId() *int64
	UserId() *int64
	Message() *string
	SetAccountId(accountId *int64)
	SetChannelId(channelId *int64)
	SetUserId(userId *int64)
	SetMessage(message *string)
	BeforeRead(ctx context.Context) error
	BeforeAdd(ctx context.Context) error
}
type messagesInterface interface {
	Len() int
	Append(req messageInterface)
	Elements() []messageInterface
}

func (m *messagesStruct) Len() int {
	if m != nil {
		return len(*m)
	}
	return 0
}
func (m *messagesStruct) Append(req messageInterface) {
	if m != nil {
		*m = append(*m, req)
	}
}
func (m *messagesStruct) Elements() []messageInterface {
	return *m
}
func (m *messageStruct) AccountId() *int64 {
	return m.accountId
}

func (m *messageStruct) ChannelId() *int64 {
	return m.channelId
}

func (m *messageStruct) UserId() *int64 {
	return m.userId
}

func (m *messageStruct) Message() *string {
	return m.message
}

func (m *messageStruct) SetAccountId(accountId *int64) {
	m.accountId = accountId
}

func (m *messageStruct) SetChannelId(channelId *int64) {
	m.channelId = channelId
}

func (m *messageStruct) SetUserId(userId *int64) {
	m.userId = userId
}

func (m *messageStruct) SetMessage(message *string) {
	m.message = message
}

func (m *messageStruct) BeforeRead(ctx context.Context) error {
	return nil
}

func (m *messageStruct) BeforeAdd(ctx context.Context) error {
	return nil
}
