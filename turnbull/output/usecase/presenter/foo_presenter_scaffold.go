package presenter

import (
	"context"
	entity "github.com/brianshepanek/turnbull/turnbull/output/domain/entity"
)

type fooPresenter interface {
	Browse(ctx context.Context, req entity.Foos) (entity.Foos, error)
	Read(ctx context.Context, req entity.Foo) (entity.Foo, error)
	Edit(ctx context.Context, req entity.Foo) (entity.Foo, error)
	Add(ctx context.Context, req entity.Foo) (entity.Foo, error)
	Delete(ctx context.Context, req entity.Foo) (entity.Foo, error)
}
