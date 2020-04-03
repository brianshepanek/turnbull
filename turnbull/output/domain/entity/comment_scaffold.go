package entity

import (
	"context"
	"time"
)

type commentStruct struct {
	Id       *int64
	PostId   *int64
	Title    *string
	Body     *string
	Created  *time.Time
	Modified *time.Time
}

func (m *commentStruct) BeforeRead(ctx context.Context) error {
	return nil
}

func (m *commentStruct) BeforeAdd(ctx context.Context) error {
	return nil
}

func (m *commentStruct) ToPrimary(ctx context.Context, req interface{}) (int64, error) {
	var resp int64
	return resp, nil
}
