package presenter

import (
	"context"
	entity "github.com/brianshepanek/turnbull/turnbull/output/domain/entity"
)

type postPresenter interface {
	Browse(ctx context.Context, req entity.Posts) (entity.Posts, error)
	Read(ctx context.Context, req entity.Post) (entity.Post, error)
	Edit(ctx context.Context, req entity.Post) (entity.Post, error)
	Add(ctx context.Context, req entity.Post) (entity.Post, error)
	Delete(ctx context.Context, req entity.Post) (entity.Post, error)
}
