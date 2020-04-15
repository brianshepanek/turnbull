package presenter

import (
	"context"
	entity "github.com/brianshepanek/turnbull/turnbull/output/domain/entity"
)

type defaultUserPresenterStruct struct{}

func (r *defaultUserPresenterStruct) Browse(ctx context.Context, req entity.Users) (entity.Users, error) {
	return req, nil
}
func (r *defaultUserPresenterStruct) BrowseByAccountId(ctx context.Context, req entity.Users) (entity.Users, error) {
	return req, nil
}
func (r *defaultUserPresenterStruct) Read(ctx context.Context, req entity.User) (entity.User, error) {
	return req, nil
}
func (r *defaultUserPresenterStruct) ReadByAccountIdAndEmail(ctx context.Context, req entity.User) (entity.User, error) {
	return req, nil
}
func (r *defaultUserPresenterStruct) Edit(ctx context.Context, req entity.User) (entity.User, error) {
	return req, nil
}
func (r *defaultUserPresenterStruct) Add(ctx context.Context, req entity.User) (entity.User, error) {
	return req, nil
}
func (r *defaultUserPresenterStruct) Delete(ctx context.Context, req entity.User) (entity.User, error) {
	return req, nil
}
