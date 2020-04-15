package interactor

import (
	"context"
	entity "github.com/brianshepanek/turnbull/turnbull/output/domain/entity"
	presenter "github.com/brianshepanek/turnbull/turnbull/output/usecase/presenter"
	repository "github.com/brianshepanek/turnbull/turnbull/output/usecase/repository"
)

type messageInteractorStruct struct {
	repository        repository.MessageRepository
	presenter         presenter.MessagePresenter
	accountInteractor AccountInteractor
	channelInteractor ChannelInteractor
	userInteractor    UserInteractor
}
type messageInteractorInterface interface {
	Browse(ctx context.Context, req entity.Messages) (entity.Messages, error)
	BrowseByAccountIdChannelId(ctx context.Context, account_id int64, channel_id int64, req entity.Messages) (entity.Messages, error)
	Read(ctx context.Context, id int64, req entity.Message) (entity.Message, error)
	Edit(ctx context.Context, id int64, req entity.Message) (entity.Message, error)
	Add(ctx context.Context, req entity.Message) (entity.Message, error)
	Delete(ctx context.Context, id int64, req entity.Message) (entity.Message, error)
}

func (i *messageInteractorStruct) Browse(ctx context.Context, req entity.Messages) (entity.Messages, error) {
	var err error
	err = i.repository.Browse(ctx, req)
	if err != nil {
		return nil, err
	}
	return i.presenter.Browse(ctx, req)
}
func (i *messageInteractorStruct) BrowseByAccountIdChannelId(ctx context.Context, account_id int64, channel_id int64, req entity.Messages) (entity.Messages, error) {
	var err error
	err = i.repository.BrowseByAccountIdChannelId(ctx, account_id, channel_id, req)
	if err != nil {
		return nil, err
	}
	return i.presenter.BrowseByAccountIdChannelId(ctx, req)
}
func (i *messageInteractorStruct) Read(ctx context.Context, id int64, req entity.Message) (entity.Message, error) {
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
func (i *messageInteractorStruct) Edit(ctx context.Context, id int64, req entity.Message) (entity.Message, error) {
	var err error
	err = i.repository.Edit(ctx, id, req)
	if err != nil {
		return nil, err
	}
	return i.presenter.Edit(ctx, req)
}
func (i *messageInteractorStruct) Add(ctx context.Context, req entity.Message) (entity.Message, error) {
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
func (i *messageInteractorStruct) Delete(ctx context.Context, id int64, req entity.Message) (entity.Message, error) {
	var err error
	err = i.repository.Delete(ctx, id, req)
	if err != nil {
		return nil, err
	}
	return i.presenter.Delete(ctx, req)
}
