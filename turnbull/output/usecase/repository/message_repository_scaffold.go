package repository

import (
	"context"
	entity "github.com/brianshepanek/turnbull/turnbull/output/domain/entity"
)

type messageRepository interface {
	Browse(ctx context.Context, req entity.Messages) error
	BrowseByAccountIdChannelId(ctx context.Context, account_id int64, channel_id int64, req entity.Messages) error
	Read(ctx context.Context, id int64, req entity.Message) error
	Edit(ctx context.Context, id int64, req entity.Message) error
	Add(ctx context.Context, req entity.Message) error
	Delete(ctx context.Context, id int64, req entity.Message) error
}
