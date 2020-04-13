package presenter

import (
	"context"
	entity "github.com/brianshepanek/turnbull/turnbull/output/domain/entity"
)

type accountPresenter interface {
	Browse(ctx context.Context, req entity.Accounts) (entity.Accounts, error)
	ReadByAccountId(ctx context.Context, req entity.Account) (entity.Account, error)
	Read(ctx context.Context, req entity.Account) (entity.Account, error)
	Edit(ctx context.Context, req entity.Account) (entity.Account, error)
	Add(ctx context.Context, req entity.Account) (entity.Account, error)
	Delete(ctx context.Context, req entity.Account) (entity.Account, error)
}
