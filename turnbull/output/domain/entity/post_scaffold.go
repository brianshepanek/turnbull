package entity

import (
	"context"
	"time"
)

type postStruct struct {
	Id       *int64
	Title    *string
	Subtitle *string
	Views    *int
	Tags     *[]string
	Created  *time.Time
	Modified *time.Time
}

func (m *postStruct) BeforeRead(ctx context.Context) error {
	return nil
}

func (m *postStruct) BeforeAdd(ctx context.Context) error {
	return nil
}

func (m *postStruct) ToPrimary(ctx context.Context, req interface{}) (int64, error) {
	var resp int64
	return resp, nil
}
