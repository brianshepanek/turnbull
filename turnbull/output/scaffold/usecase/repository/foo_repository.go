package repository

import (
	"context"
	entity "github.com/brianshepanek/turnbull/turnbull/output/scaffold/domain/entity"
)

type FooScaffoldRepository interface {
	Count(ctx context.Context, query interface{}, req *int) error
	Browse(ctx context.Context, query interface{}, req entity.FoosScaffoldInterface) error
	Read(ctx context.Context, query interface{}, req entity.FooScaffoldInterface) error
	Edit(ctx context.Context, req entity.FooScaffoldInterface) error
	Add(ctx context.Context, req entity.FooScaffoldInterface) error
	Delete(ctx context.Context, req entity.FooScaffoldInterface) error
}
