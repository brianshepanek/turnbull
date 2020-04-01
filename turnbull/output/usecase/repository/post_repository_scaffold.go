package repository

import (
	"context"
	entity "github.com/brianshepanek/turnbull/turnbull/output/domain/entity"
)

type postRepository interface {
	Browse(ctx context.Context, req entity.Posts) error
	Read(ctx context.Context, id int64, req entity.Post) error
	Edit(ctx context.Context, id int64, req entity.Post) error
	Add(ctx context.Context, req entity.Post) error
	Delete(ctx context.Context, id int64, req entity.Post) error
}
