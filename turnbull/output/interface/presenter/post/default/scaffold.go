package presenter

import (
	"context"
	entity "github.com/brianshepanek/turnbull/turnbull/output/domain/entity"
)

type defaultPostPresenterStruct struct{}

func (r *defaultPostPresenterStruct) Browse(ctx context.Context, req *[]entity.Post) (*[]entity.Post, error) {
	return req, nil
}
func (r *defaultPostPresenterStruct) Read(ctx context.Context, req *entity.Post) (*entity.Post, error) {
	return req, nil
}
func (r *defaultPostPresenterStruct) Edit(ctx context.Context, req *entity.Post) (*entity.Post, error) {
	return req, nil
}
func (r *defaultPostPresenterStruct) Add(ctx context.Context, req *entity.Post) (*entity.Post, error) {
	return req, nil
}
func (r *defaultPostPresenterStruct) Delete(ctx context.Context, req *entity.Post) (*entity.Post, error) {
	return req, nil
}
