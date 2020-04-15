package repository

import (
	"context"
	entity "github.com/brianshepanek/turnbull/turnbull/output/domain/entity"
)

type channelRepository interface {
	Browse(ctx context.Context, req entity.Channels) error
	BrowseByAccountId(ctx context.Context, account_id int64, req entity.Channels) error
	Read(ctx context.Context, id int64, req entity.Channel) error
	ReadByAccountIdAndName(ctx context.Context, account_id int64, name string, req entity.Channel) error
	Edit(ctx context.Context, id int64, req entity.Channel) error
	Add(ctx context.Context, req entity.Channel) error
	Delete(ctx context.Context, id int64, req entity.Channel) error
}
