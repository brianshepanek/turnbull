package interactor

import (
	"context"
	entity "github.com/brianshepanek/turnbull/turnbull/output/domain/entity"
	presenter "github.com/brianshepanek/turnbull/turnbull/output/usecase/presenter"
	repository "github.com/brianshepanek/turnbull/turnbull/output/usecase/repository"
)

type accountInteractorStruct struct {
	repository repository.AccountRepository
	presenter  presenter.AccountPresenter
}
type accountInteractorInterface interface {
	Browse(ctx context.Context, req entity.Accounts) (entity.Accounts, error)
	Read(ctx context.Context, id int64, req entity.Account) (entity.Account, error)
	Edit(ctx context.Context, id int64, req entity.Account) (entity.Account, error)
	Add(ctx context.Context, req entity.Account) (entity.Account, error)
	Delete(ctx context.Context, id int64, req entity.Account) (entity.Account, error)
}

func (i *accountInteractorStruct) Browse(ctx context.Context, req entity.Accounts) (entity.Accounts, error) {
	var err error
	err = i.repository.Browse(ctx, req)
	if err != nil {
		return nil, err
	}
	return i.presenter.Browse(ctx, req)
}
func (i *accountInteractorStruct) Read(ctx context.Context, id int64, req entity.Account) (entity.Account, error) {
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
func (i *accountInteractorStruct) Edit(ctx context.Context, id int64, req entity.Account) (entity.Account, error) {
	var err error
	err = i.repository.Edit(ctx, id, req)
	if err != nil {
		return nil, err
	}
	return i.presenter.Edit(ctx, req)
}
func (i *accountInteractorStruct) Add(ctx context.Context, req entity.Account) (entity.Account, error) {
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
func (i *accountInteractorStruct) Delete(ctx context.Context, id int64, req entity.Account) (entity.Account, error) {
	var err error
	err = i.repository.Delete(ctx, id, req)
	if err != nil {
		return nil, err
	}
	return i.presenter.Delete(ctx, req)
}
