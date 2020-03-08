package interactor

import (
	"context"
	model "github.com/brianshepanek/turnbull/testing/output/scaffold/domain/model"
	presenter "github.com/brianshepanek/turnbull/testing/output/scaffold/usecase/presenter"
	repository "github.com/brianshepanek/turnbull/testing/output/scaffold/usecase/repository"
)

type fooScaffoldInteractor struct {
	FooScaffoldRepository repository.FooScaffoldRepository
	FooScaffoldPresenter  presenter.FooScaffoldPresenter
}

type FooScaffoldInteractor interface {
	Add(ctx context.Context, req model.FooScaffoldInterface) (*model.FooScaffoldInterface, error)
}

func NewFooScaffoldInteractor(r repository.FooScaffoldRepository, p presenter.FooScaffoldPresenter) FooScaffoldInteractor {
	return &fooScaffoldInteractor{r, p}
}

func (i *fooScaffoldInteractor) Add(ctx context.Context, req model.FooScaffoldInterface) (*model.FooScaffoldInterface, error) {
	err := i.FooScaffoldRepository.Add(ctx, req)
	if err != nil {
		return nil, err
	}
	return i.FooScaffoldPresenter.Add(ctx, req)
}
