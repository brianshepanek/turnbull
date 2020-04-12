package repository

import (
	"context"
	entity "github.com/brianshepanek/turnbull/turnbull/output/domain/entity"
)

type enhancedAccountRepository interface {
	Browse(ctx context.Context, req entity.EnhancedAccounts) error
	Read(ctx context.Context, id int64, req entity.EnhancedAccount) error
	Edit(ctx context.Context, id int64, req entity.EnhancedAccount) error
	Add(ctx context.Context, req entity.EnhancedAccount) error
	Delete(ctx context.Context, id int64, req entity.EnhancedAccount) error
}
