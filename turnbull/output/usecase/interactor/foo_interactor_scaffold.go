package interactor

import (
	"context"
	entity "github.com/brianshepanek/turnbull/turnbull/output/domain/entity"
	presenter "github.com/brianshepanek/turnbull/turnbull/output/usecase/presenter"
	repository "github.com/brianshepanek/turnbull/turnbull/output/usecase/repository"
)

type fooInteractorStruct struct {
	repository repository.FooRepository
	presenter  presenter.FooPresenter
}
type fooInteractorInterface interface {
	Browse(ctx context.Context, req entity.Foos) (entity.Foos, error)
	Read(ctx context.Context, id int64, req entity.Foo) (entity.Foo, error)
	Edit(ctx context.Context, id int64, req entity.Foo) (entity.Foo, error)
	Add(ctx context.Context, req entity.Foo) (entity.Foo, error)
	Delete(ctx context.Context, id int64, req entity.Foo) (entity.Foo, error)
}

func (i *fooInteractorStruct) Browse(ctx context.Context, req entity.Foos) (entity.Foos, error) {
	var err error
	err = i.repository.Browse(ctx, req)
	if err != nil {
		return nil, err
	}
	return i.presenter.Browse(ctx, req)
}
func (i *fooInteractorStruct) Read(ctx context.Context, id int64, req entity.Foo) (entity.Foo, error) {
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
func (i *fooInteractorStruct) Edit(ctx context.Context, id int64, req entity.Foo) (entity.Foo, error) {
	var err error
	err = i.repository.Edit(ctx, id, req)
	if err != nil {
		return nil, err
	}
	return i.presenter.Edit(ctx, req)
}
func (i *fooInteractorStruct) Add(ctx context.Context, req entity.Foo) (entity.Foo, error) {
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
func (i *fooInteractorStruct) Delete(ctx context.Context, id int64, req entity.Foo) (entity.Foo, error) {
	var err error
	err = i.repository.Delete(ctx, id, req)
	if err != nil {
		return nil, err
	}
	return i.presenter.Delete(ctx, req)
}
