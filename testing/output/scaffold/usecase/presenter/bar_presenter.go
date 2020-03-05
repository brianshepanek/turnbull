package presenter

import (
	"context"
	model "github.com/brianshepanek/turnbull/testing/output/scaffold/domain/model"
)

type BarScaffoldPresenter interface {
	Add(ctx context.Context, req model.BarScaffoldInterface) (*model.BarScaffoldInterface, error)
	FindAll(ctx context.Context, req []model.BarScaffoldInterface) (*[]model.BarScaffoldInterface, error)
	FindOne(ctx context.Context, req model.BarScaffoldInterface) (*model.BarScaffoldInterface, error)
	CountEmUp(ctx context.Context, req int) (*int, error)
}
