package presenter

import (
	"context"
	entity "github.com/brianshepanek/turnbull/turnbull/output/domain/entity"
)

type userPresenter interface {
	Browse(ctx context.Context, req entity.Users) (entity.Users, error)
	BrowseByAccountId(ctx context.Context, req entity.Users) (entity.Users, error)
	Read(ctx context.Context, req entity.User) (entity.User, error)
	Edit(ctx context.Context, req entity.User) (entity.User, error)
	Add(ctx context.Context, req entity.User) (entity.User, error)
	Delete(ctx context.Context, req entity.User) (entity.User, error)
}
