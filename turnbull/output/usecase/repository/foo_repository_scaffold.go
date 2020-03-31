package repository

import (
	"context"
	entity "github.com/brianshepanek/turnbull/turnbull/output/domain/entity"
)

type fooRepository interface {
	Browse(ctx context.Context, req entity.Foos) error
	Read(ctx context.Context, id int64, req entity.Foo) error
	Edit(ctx context.Context, id int64, req entity.Foo) error
	Add(ctx context.Context, req entity.Foo) error
	Delete(ctx context.Context, id int64, req entity.Foo) error
}
