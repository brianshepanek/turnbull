package presenter

import (
	"context"
	entity "github.com/brianshepanek/turnbull/turnbull/output/scaffold/domain/entity"
)

type FooScaffoldPresenter interface {
	Count(ctx context.Context, req *int) (*int, error)
	Browse(ctx context.Context, req entity.FoosScaffoldInterface) (*entity.FoosScaffoldInterface, error)
	Read(ctx context.Context, req entity.FooScaffoldInterface) (*entity.FooScaffoldInterface, error)
	Edit(ctx context.Context, req entity.FooScaffoldInterface) (*entity.FooScaffoldInterface, error)
	Add(ctx context.Context, req entity.FooScaffoldInterface) (*entity.FooScaffoldInterface, error)
	Delete(ctx context.Context, req entity.FooScaffoldInterface) (*entity.FooScaffoldInterface, error)
}
