package presenter

import (
	"context"
	model "github.com/brianshepanek/turnbull/testing/output/scaffold/domain/model"
)

type FooScaffoldPresenter interface {
	Add(ctx context.Context, req model.FooScaffoldInterface) (*model.FooScaffoldInterface, error)
}
