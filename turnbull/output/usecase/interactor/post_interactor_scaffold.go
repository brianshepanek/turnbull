package interactor

import (
	"context"
	entity "github.com/brianshepanek/turnbull/turnbull/output/domain/entity"
	presenter "github.com/brianshepanek/turnbull/turnbull/output/usecase/presenter"
	repository "github.com/brianshepanek/turnbull/turnbull/output/usecase/repository"
)

type postInteractorStruct struct {
	repository repository.PostRepository
	presenter  presenter.PostPresenter
}
type postInteractorInterface interface {
	Browse(ctx context.Context, req entity.Posts) (entity.Posts, error)
	Read(ctx context.Context, id int64, req entity.Post) (entity.Post, error)
	Edit(ctx context.Context, id int64, req entity.Post) (entity.Post, error)
	Add(ctx context.Context, req entity.Post) (entity.Post, error)
	Delete(ctx context.Context, id int64, req entity.Post) (entity.Post, error)
}

func (i *postInteractorStruct) Browse(ctx context.Context, req entity.Posts) (entity.Posts, error) {
	var err error
	err = i.repository.Browse(ctx, req)
	if err != nil {
		return nil, err
	}
	return i.presenter.Browse(ctx, req)
}
func (i *postInteractorStruct) Read(ctx context.Context, id int64, req entity.Post) (entity.Post, error) {
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
func (i *postInteractorStruct) Edit(ctx context.Context, id int64, req entity.Post) (entity.Post, error) {
	var err error
	err = i.repository.Edit(ctx, id, req)
	if err != nil {
		return nil, err
	}
	return i.presenter.Edit(ctx, req)
}
func (i *postInteractorStruct) Add(ctx context.Context, req entity.Post) (entity.Post, error) {
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
func (i *postInteractorStruct) Delete(ctx context.Context, id int64, req entity.Post) (entity.Post, error) {
	var err error
	err = i.repository.Delete(ctx, id, req)
	if err != nil {
		return nil, err
	}
	return i.presenter.Delete(ctx, req)
}
