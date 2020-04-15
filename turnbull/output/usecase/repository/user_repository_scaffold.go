package repository

import (
	"context"
	entity "github.com/brianshepanek/turnbull/turnbull/output/domain/entity"
)

type userRepository interface {
	Browse(ctx context.Context, req entity.Users) error
	BrowseByAccountId(ctx context.Context, account_id int64, req entity.Users) error
	Read(ctx context.Context, id int64, req entity.User) error
	ReadByAccountIdAndEmail(ctx context.Context, account_id int64, email string, req entity.User) error
	Edit(ctx context.Context, id int64, req entity.User) error
	Add(ctx context.Context, req entity.User) error
	Delete(ctx context.Context, id int64, req entity.User) error
}
