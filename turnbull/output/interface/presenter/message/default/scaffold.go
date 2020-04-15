package presenter

import (
	"context"
	entity "github.com/brianshepanek/turnbull/turnbull/output/domain/entity"
)

type defaultMessagePresenterStruct struct{}

func (r *defaultMessagePresenterStruct) Browse(ctx context.Context, req entity.Messages) (entity.Messages, error) {
	return req, nil
}
func (r *defaultMessagePresenterStruct) BrowseByAccountIdChannelId(ctx context.Context, req entity.Messages) (entity.Messages, error) {
	return req, nil
}
func (r *defaultMessagePresenterStruct) Read(ctx context.Context, req entity.Message) (entity.Message, error) {
	return req, nil
}
func (r *defaultMessagePresenterStruct) Edit(ctx context.Context, req entity.Message) (entity.Message, error) {
	return req, nil
}
func (r *defaultMessagePresenterStruct) Add(ctx context.Context, req entity.Message) (entity.Message, error) {
	return req, nil
}
func (r *defaultMessagePresenterStruct) Delete(ctx context.Context, req entity.Message) (entity.Message, error) {
	return req, nil
}
