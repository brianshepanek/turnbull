package presenter

import (
	"context"
	entity "github.com/brianshepanek/turnbull/turnbull/output/domain/entity"
)

type defaultEnhancedAccountPresenterStruct struct{}

func (r *defaultEnhancedAccountPresenterStruct) Browse(ctx context.Context, req entity.EnhancedAccounts) (entity.EnhancedAccounts, error) {
	return req, nil
}
func (r *defaultEnhancedAccountPresenterStruct) Read(ctx context.Context, req entity.EnhancedAccount) (entity.EnhancedAccount, error) {
	return req, nil
}
func (r *defaultEnhancedAccountPresenterStruct) Edit(ctx context.Context, req entity.EnhancedAccount) (entity.EnhancedAccount, error) {
	return req, nil
}
func (r *defaultEnhancedAccountPresenterStruct) Add(ctx context.Context, req entity.EnhancedAccount) (entity.EnhancedAccount, error) {
	return req, nil
}
func (r *defaultEnhancedAccountPresenterStruct) Delete(ctx context.Context, req entity.EnhancedAccount) (entity.EnhancedAccount, error) {
	return req, nil
}
