package interactor

import (
	"context"
	model "github.com/brianshepanek/turnbull/testing/output/scaffold/domain/model"
	presenter "github.com/brianshepanek/turnbull/testing/output/scaffold/usecase/presenter"
	repository "github.com/brianshepanek/turnbull/testing/output/scaffold/usecase/repository"
)

type barScaffoldInteractor struct {
	BarScaffoldRepository repository.BarScaffoldRepository
	BarScaffoldPresenter  presenter.BarScaffoldPresenter
}

type BarScaffoldInteractor interface {
	Add(ctx context.Context, req model.BarScaffoldInterface) (*model.BarScaffoldInterface, error)
	FindAll(ctx context.Context, req interface{}, resp model.BarsScaffoldInterface) (*model.BarsScaffoldInterface, error)
	FindOne(ctx context.Context, req model.BarScaffoldInterface) (*model.BarScaffoldInterface, error)
	CountEmUp(ctx context.Context, req interface{}) (*int, error)
	ByeBye(ctx context.Context, req model.BarScaffoldInterface) (*model.BarScaffoldInterface, error)
}

func NewBarScaffoldInteractor(r repository.BarScaffoldRepository, p presenter.BarScaffoldPresenter) BarScaffoldInteractor {
	return &barScaffoldInteractor{r, p}
}

func (i *barScaffoldInteractor) Add(ctx context.Context, req model.BarScaffoldInterface) (*model.BarScaffoldInterface, error) {
	err := i.BarScaffoldRepository.Add(ctx, req)
	if err != nil {
		return nil, err
	}
	return i.BarScaffoldPresenter.Add(ctx, req)
}

func (i *barScaffoldInteractor) FindAll(ctx context.Context, req interface{}, resp model.BarsScaffoldInterface) (*model.BarsScaffoldInterface, error) {
	err := i.BarScaffoldRepository.FindAll(ctx, req, resp)
	if err != nil {
		return nil, err
	}
	return i.BarScaffoldPresenter.FindAll(ctx, req)
}

func (i *barScaffoldInteractor) FindOne(ctx context.Context, req model.BarScaffoldInterface) (*model.BarScaffoldInterface, error) {
	err := i.BarScaffoldRepository.FindOne(ctx, req)
	if err != nil {
		return nil, err
	}
	return i.BarScaffoldPresenter.FindOne(ctx, req)
}

func (i *barScaffoldInteractor) CountEmUp(ctx context.Context, req interface{}) (*int, error) {
	resp, err := i.BarScaffoldRepository.CountEmUp(ctx, req)
	if err != nil {
		return nil, err
	}
	return i.BarScaffoldPresenter.CountEmUp(ctx, *resp)
}

func (i *barScaffoldInteractor) ByeBye(ctx context.Context, req model.BarScaffoldInterface) (*model.BarScaffoldInterface, error) {
	err := i.BarScaffoldRepository.ByeBye(ctx, req)
	if err != nil {
		return nil, err
	}
	return i.BarScaffoldPresenter.ByeBye(ctx, req)
}
