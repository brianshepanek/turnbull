package registry

import (
	default1 "github.com/brianshepanek/turnbull/turnbull/output/interface/presenter/enhanced_account/default"
	presenter "github.com/brianshepanek/turnbull/turnbull/output/usecase/presenter"
)

type enhancedAccountDefaultPresenterRegistry struct{}

func (r *registry) RegisterDefaultEnhancedAccountPresenter() {}
func (r *registry) newDefaultEnhancedAccountPresenter() presenter.EnhancedAccountPresenter {
	return default1.New()
}
