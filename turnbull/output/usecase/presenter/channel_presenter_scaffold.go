package presenter

import (
	"context"
	entity "github.com/brianshepanek/turnbull/turnbull/output/domain/entity"
)

type channelPresenter interface {
	Browse(ctx context.Context, req entity.Channels) (entity.Channels, error)
	BrowseByAccountId(ctx context.Context, req entity.Channels) (entity.Channels, error)
	Read(ctx context.Context, req entity.Channel) (entity.Channel, error)
	ReadByAccountIdAndName(ctx context.Context, req entity.Channel) (entity.Channel, error)
	Edit(ctx context.Context, req entity.Channel) (entity.Channel, error)
	Add(ctx context.Context, req entity.Channel) (entity.Channel, error)
	Delete(ctx context.Context, req entity.Channel) (entity.Channel, error)
}
