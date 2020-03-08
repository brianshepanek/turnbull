package repository

import (
	"context"
	model "github.com/brianshepanek/turnbull/testing/output/scaffold/domain/model"
)

type BarScaffoldRepository interface {
	Add(ctx context.Context, req model.BarScaffoldInterface) (err error)
	FindAll(ctx context.Context, req interface{}, resp model.BarsScaffoldInterface) (err error)
	FindOne(ctx context.Context, req model.BarScaffoldInterface) (err error)
	CountEmUp(ctx context.Context, req interface{}) (resp *int, err error)
	ByeBye(ctx context.Context, req model.BarScaffoldInterface) (err error)
}
