package repository

import (
	"context"
	entity "github.com/brianshepanek/turnbull/turnbull/output/domain/entity"
)

type commentRepository interface {
	Browse(ctx context.Context, req *[]entity.Comment) error
	Read(ctx context.Context, id int64, req *entity.Comment) error
	Edit(ctx context.Context, id int64, req *entity.Comment) error
	Add(ctx context.Context, req *entity.Comment) error
	Delete(ctx context.Context, id int64, req *entity.Comment) error
}
