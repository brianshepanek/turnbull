package repository

import (
	"context"
	entity "github.com/brianshepanek/turnbull/turnbull/output/domain/entity"
)

type accountRepository interface {
	Browse(ctx context.Context, req entity.Accounts) error
	ReadByAccountId(ctx context.Context, id int64, req entity.Account) error
	Read(ctx context.Context, id int64, req entity.Account) error
	Edit(ctx context.Context, id int64, req entity.Account) error
	Add(ctx context.Context, req entity.Account) error
	Delete(ctx context.Context, id int64, req entity.Account) error
}
