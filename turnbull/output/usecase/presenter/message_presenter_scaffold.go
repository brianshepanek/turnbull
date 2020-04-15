package presenter

import (
	"context"
	entity "github.com/brianshepanek/turnbull/turnbull/output/domain/entity"
)

type messagePresenter interface {
	Browse(ctx context.Context, req entity.Messages) (entity.Messages, error)
	BrowseByAccountIdChannelId(ctx context.Context, req entity.Messages) (entity.Messages, error)
	Read(ctx context.Context, req entity.Message) (entity.Message, error)
	Edit(ctx context.Context, req entity.Message) (entity.Message, error)
	Add(ctx context.Context, req entity.Message) (entity.Message, error)
	Delete(ctx context.Context, req entity.Message) (entity.Message, error)
}
