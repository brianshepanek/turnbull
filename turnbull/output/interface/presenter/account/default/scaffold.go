package presenter

import (
	"context"
	entity "github.com/brianshepanek/turnbull/turnbull/output/domain/entity"
)

type defaultAccountPresenterStruct struct{}

func (r *defaultAccountPresenterStruct) Browse(ctx context.Context, req entity.Accounts) (entity.Accounts, error) {
	return req, nil
}
func (r *defaultAccountPresenterStruct) Read(ctx context.Context, req entity.Account) (entity.Account, error) {
	return req, nil
}
func (r *defaultAccountPresenterStruct) Edit(ctx context.Context, req entity.Account) (entity.Account, error) {
	return req, nil
}
func (r *defaultAccountPresenterStruct) Add(ctx context.Context, req entity.Account) (entity.Account, error) {
	return req, nil
}
func (r *defaultAccountPresenterStruct) Delete(ctx context.Context, req entity.Account) (entity.Account, error) {
	return req, nil
}
