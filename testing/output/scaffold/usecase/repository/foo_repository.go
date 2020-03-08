package repository

import (
	"context"
	model "github.com/brianshepanek/turnbull/testing/output/scaffold/domain/model"
)

type FooScaffoldRepository interface {
	Add(ctx context.Context, req model.FooScaffoldInterface) (err error)
}
