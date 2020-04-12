package presenter

import (
	"context"
	entity "github.com/brianshepanek/turnbull/turnbull/output/domain/entity"
)

type enhancedAccountPresenter interface {
	Browse(ctx context.Context, req entity.EnhancedAccounts) (entity.EnhancedAccounts, error)
	Read(ctx context.Context, req entity.EnhancedAccount) (entity.EnhancedAccount, error)
	Edit(ctx context.Context, req entity.EnhancedAccount) (entity.EnhancedAccount, error)
	Add(ctx context.Context, req entity.EnhancedAccount) (entity.EnhancedAccount, error)
	Delete(ctx context.Context, req entity.EnhancedAccount) (entity.EnhancedAccount, error)
}
