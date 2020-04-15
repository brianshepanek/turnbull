package interactor

import (
	"context"
	entity "github.com/brianshepanek/turnbull/turnbull/output/domain/entity"
	presenter "github.com/brianshepanek/turnbull/turnbull/output/usecase/presenter"
	repository "github.com/brianshepanek/turnbull/turnbull/output/usecase/repository"
)

type channelInteractorStruct struct {
	repository        repository.ChannelRepository
	presenter         presenter.ChannelPresenter
	accountInteractor AccountInteractor
}
type channelInteractorInterface interface {
	Browse(ctx context.Context, req entity.Channels) (entity.Channels, error)
	BrowseByAccountId(ctx context.Context, account_id int64, req entity.Channels) (entity.Channels, error)
	Read(ctx context.Context, id int64, req entity.Channel) (entity.Channel, error)
	ReadByAccountIdAndName(ctx context.Context, account_id int64, name string, req entity.Channel) (entity.Channel, error)
	Edit(ctx context.Context, id int64, req entity.Channel) (entity.Channel, error)
	Add(ctx context.Context, req entity.Channel) (entity.Channel, error)
	Delete(ctx context.Context, id int64, req entity.Channel) (entity.Channel, error)
}

func (i *channelInteractorStruct) Browse(ctx context.Context, req entity.Channels) (entity.Channels, error) {
	var err error
	err = i.repository.Browse(ctx, req)
	if err != nil {
		return nil, err
	}
	return i.presenter.Browse(ctx, req)
}
func (i *channelInteractorStruct) BrowseByAccountId(ctx context.Context, account_id int64, req entity.Channels) (entity.Channels, error) {
	var err error
	err = i.repository.BrowseByAccountId(ctx, account_id, req)
	if err != nil {
		return nil, err
	}
	return i.presenter.BrowseByAccountId(ctx, req)
}
func (i *channelInteractorStruct) Read(ctx context.Context, id int64, req entity.Channel) (entity.Channel, error) {
	var err error
	err = req.BeforeRead(ctx)
	if err != nil {
		return nil, err
	}
	err = i.repository.Read(ctx, id, req)
	if err != nil {
		return nil, err
	}
	return i.presenter.Read(ctx, req)
}
func (i *channelInteractorStruct) ReadByAccountIdAndName(ctx context.Context, account_id int64, name string, req entity.Channel) (entity.Channel, error) {
	var err error
	err = i.repository.ReadByAccountIdAndName(ctx, account_id, name, req)
	if err != nil {
		return nil, err
	}
	return i.presenter.ReadByAccountIdAndName(ctx, req)
}
func (i *channelInteractorStruct) Edit(ctx context.Context, id int64, req entity.Channel) (entity.Channel, error) {
	var err error
	err = i.repository.Edit(ctx, id, req)
	if err != nil {
		return nil, err
	}
	return i.presenter.Edit(ctx, req)
}
func (i *channelInteractorStruct) Add(ctx context.Context, req entity.Channel) (entity.Channel, error) {
	var err error
	err = req.BeforeAdd(ctx)
	if err != nil {
		return nil, err
	}
	err = i.repository.Add(ctx, req)
	if err != nil {
		return nil, err
	}
	return i.presenter.Add(ctx, req)
}
func (i *channelInteractorStruct) Delete(ctx context.Context, id int64, req entity.Channel) (entity.Channel, error) {
	var err error
	err = i.repository.Delete(ctx, id, req)
	if err != nil {
		return nil, err
	}
	return i.presenter.Delete(ctx, req)
}
