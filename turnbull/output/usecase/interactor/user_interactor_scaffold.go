package interactor

import (
	"context"
	entity "github.com/brianshepanek/turnbull/turnbull/output/domain/entity"
	presenter "github.com/brianshepanek/turnbull/turnbull/output/usecase/presenter"
	repository "github.com/brianshepanek/turnbull/turnbull/output/usecase/repository"
)

type userInteractorStruct struct {
	repository repository.UserRepository
	presenter  presenter.UserPresenter
}
type userInteractorInterface interface {
	Browse(ctx context.Context, req entity.Users) (entity.Users, error)
	BrowseByAccountId(ctx context.Context, id int64, req entity.Users) (entity.Users, error)
	Read(ctx context.Context, id int64, req entity.User) (entity.User, error)
	Edit(ctx context.Context, id int64, req entity.User) (entity.User, error)
	Add(ctx context.Context, req entity.User) (entity.User, error)
	Delete(ctx context.Context, id int64, req entity.User) (entity.User, error)
}

func (i *userInteractorStruct) Browse(ctx context.Context, req entity.Users) (entity.Users, error) {
	var err error
	err = i.repository.Browse(ctx, req)
	if err != nil {
		return nil, err
	}
	return i.presenter.Browse(ctx, req)
}
func (i *userInteractorStruct) BrowseByAccountId(ctx context.Context, id int64, req entity.Users) (entity.Users, error) {
	var err error
	err = i.repository.BrowseByAccountId(ctx, id, req)
	if err != nil {
		return nil, err
	}
	return i.presenter.BrowseByAccountId(ctx, req)
}
func (i *userInteractorStruct) Read(ctx context.Context, id int64, req entity.User) (entity.User, error) {
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
func (i *userInteractorStruct) Edit(ctx context.Context, id int64, req entity.User) (entity.User, error) {
	var err error
	err = i.repository.Edit(ctx, id, req)
	if err != nil {
		return nil, err
	}
	return i.presenter.Edit(ctx, req)
}
func (i *userInteractorStruct) Add(ctx context.Context, req entity.User) (entity.User, error) {
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
func (i *userInteractorStruct) Delete(ctx context.Context, id int64, req entity.User) (entity.User, error) {
	var err error
	err = i.repository.Delete(ctx, id, req)
	if err != nil {
		return nil, err
	}
	return i.presenter.Delete(ctx, req)
}
