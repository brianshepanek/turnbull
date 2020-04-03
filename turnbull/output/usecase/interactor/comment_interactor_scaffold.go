package interactor

import (
	"context"
	entity "github.com/brianshepanek/turnbull/turnbull/output/domain/entity"
	presenter "github.com/brianshepanek/turnbull/turnbull/output/usecase/presenter"
	repository "github.com/brianshepanek/turnbull/turnbull/output/usecase/repository"
)

type commentInteractorStruct struct {
	repository repository.CommentRepository
	presenter  presenter.CommentPresenter
}
type commentInteractorInterface interface {
	Browse(ctx context.Context, req *[]entity.Comment) (*[]entity.Comment, error)
	Read(ctx context.Context, id int64, req *entity.Comment) (*entity.Comment, error)
	Edit(ctx context.Context, id int64, req *entity.Comment) (*entity.Comment, error)
	Add(ctx context.Context, req *entity.Comment) (*entity.Comment, error)
	Delete(ctx context.Context, id int64, req *entity.Comment) (*entity.Comment, error)
}

func (i *commentInteractorStruct) Browse(ctx context.Context, req *[]entity.Comment) (*[]entity.Comment, error) {
	var err error
	err = i.repository.Browse(ctx, req)
	if err != nil {
		return nil, err
	}
	return i.presenter.Browse(ctx, req)
}
func (i *commentInteractorStruct) Read(ctx context.Context, id int64, req *entity.Comment) (*entity.Comment, error) {
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
func (i *commentInteractorStruct) Edit(ctx context.Context, id int64, req *entity.Comment) (*entity.Comment, error) {
	var err error
	err = i.repository.Edit(ctx, id, req)
	if err != nil {
		return nil, err
	}
	return i.presenter.Edit(ctx, req)
}
func (i *commentInteractorStruct) Add(ctx context.Context, req *entity.Comment) (*entity.Comment, error) {
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
func (i *commentInteractorStruct) Delete(ctx context.Context, id int64, req *entity.Comment) (*entity.Comment, error) {
	var err error
	err = i.repository.Delete(ctx, id, req)
	if err != nil {
		return nil, err
	}
	return i.presenter.Delete(ctx, req)
}
