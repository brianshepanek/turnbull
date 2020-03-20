package repository

import (
	"context"
	entity "github.com/brianshepanek/turnbull/turnbull/output/domain/entity"
)

type fooRepository interface {
	Browse(ctx context.Context, query interface{}, req entity.Foos) error
	Read(ctx context.Context, query interface{}, req entity.Foo) error
	Edit(ctx context.Context, req entity.Foo) error
	Add(ctx context.Context, req entity.Foo) error
	Delete(ctx context.Context, req entity.Foo) error
}
