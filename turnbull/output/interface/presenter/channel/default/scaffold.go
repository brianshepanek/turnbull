package presenter

import (
	"context"
	entity "github.com/brianshepanek/turnbull/turnbull/output/domain/entity"
)

type defaultChannelPresenterStruct struct{}

func (r *defaultChannelPresenterStruct) Browse(ctx context.Context, req entity.Channels) (entity.Channels, error) {
	return req, nil
}
func (r *defaultChannelPresenterStruct) BrowseByAccountId(ctx context.Context, req entity.Channels) (entity.Channels, error) {
	return req, nil
}
func (r *defaultChannelPresenterStruct) Read(ctx context.Context, req entity.Channel) (entity.Channel, error) {
	return req, nil
}
func (r *defaultChannelPresenterStruct) ReadByAccountIdAndName(ctx context.Context, req entity.Channel) (entity.Channel, error) {
	return req, nil
}
func (r *defaultChannelPresenterStruct) Edit(ctx context.Context, req entity.Channel) (entity.Channel, error) {
	return req, nil
}
func (r *defaultChannelPresenterStruct) Add(ctx context.Context, req entity.Channel) (entity.Channel, error) {
	return req, nil
}
func (r *defaultChannelPresenterStruct) Delete(ctx context.Context, req entity.Channel) (entity.Channel, error) {
	return req, nil
}
