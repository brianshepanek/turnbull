package presenter

import (
	"context"
	entity "github.com/brianshepanek/turnbull/turnbull/output/domain/entity"
)

type commentPresenter interface {
	Browse(ctx context.Context, req *[]entity.Comment) (*[]entity.Comment, error)
	Read(ctx context.Context, req *entity.Comment) (*entity.Comment, error)
	Edit(ctx context.Context, req *entity.Comment) (*entity.Comment, error)
	Add(ctx context.Context, req *entity.Comment) (*entity.Comment, error)
	Delete(ctx context.Context, req *entity.Comment) (*entity.Comment, error)
}
