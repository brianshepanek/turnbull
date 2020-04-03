package presenter

import (
	"context"
	entity "github.com/brianshepanek/turnbull/turnbull/output/domain/entity"
)

type defaultCommentPresenterStruct struct{}

func (r *defaultCommentPresenterStruct) Browse(ctx context.Context, req *[]entity.Comment) (*[]entity.Comment, error) {
	return req, nil
}
func (r *defaultCommentPresenterStruct) Read(ctx context.Context, req *entity.Comment) (*entity.Comment, error) {
	return req, nil
}
func (r *defaultCommentPresenterStruct) Edit(ctx context.Context, req *entity.Comment) (*entity.Comment, error) {
	return req, nil
}
func (r *defaultCommentPresenterStruct) Add(ctx context.Context, req *entity.Comment) (*entity.Comment, error) {
	return req, nil
}
func (r *defaultCommentPresenterStruct) Delete(ctx context.Context, req *entity.Comment) (*entity.Comment, error) {
	return req, nil
}
