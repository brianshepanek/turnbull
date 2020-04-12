package interactor

import (
	"context"
	entity "github.com/brianshepanek/turnbull/turnbull/output/domain/entity"
	presenter "github.com/brianshepanek/turnbull/turnbull/output/usecase/presenter"
	repository "github.com/brianshepanek/turnbull/turnbull/output/usecase/repository"
)

type enhancedAccountInteractorStruct struct {
	repository repository.EnhancedAccountRepository
	presenter  presenter.EnhancedAccountPresenter
}
type enhancedAccountInteractorInterface interface {
	Browse(ctx context.Context, req entity.EnhancedAccounts) (entity.EnhancedAccounts, error)
	Read(ctx context.Context, id int64, req entity.EnhancedAccount) (entity.EnhancedAccount, error)
	Edit(ctx context.Context, id int64, req entity.EnhancedAccount) (entity.EnhancedAccount, error)
	Add(ctx context.Context, req entity.EnhancedAccount) (entity.EnhancedAccount, error)
	Delete(ctx context.Context, id int64, req entity.EnhancedAccount) (entity.EnhancedAccount, error)
}

func (i *enhancedAccountInteractorStruct) Browse(ctx context.Context, req entity.EnhancedAccounts) (entity.EnhancedAccounts, error) {
	var err error
	err = i.repository.Browse(ctx, req)
	if err != nil {
		return nil, err
	}
	return i.presenter.Browse(ctx, req)
}
func (i *enhancedAccountInteractorStruct) Read(ctx context.Context, id int64, req entity.EnhancedAccount) (entity.EnhancedAccount, error) {
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
func (i *enhancedAccountInteractorStruct) Edit(ctx context.Context, id int64, req entity.EnhancedAccount) (entity.EnhancedAccount, error) {
	var err error
	err = i.repository.Edit(ctx, id, req)
	if err != nil {
		return nil, err
	}
	return i.presenter.Edit(ctx, req)
}
func (i *enhancedAccountInteractorStruct) Add(ctx context.Context, req entity.EnhancedAccount) (entity.EnhancedAccount, error) {
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
func (i *enhancedAccountInteractorStruct) Delete(ctx context.Context, id int64, req entity.EnhancedAccount) (entity.EnhancedAccount, error) {
	var err error
	err = i.repository.Delete(ctx, id, req)
	if err != nil {
		return nil, err
	}
	return i.presenter.Delete(ctx, req)
}
