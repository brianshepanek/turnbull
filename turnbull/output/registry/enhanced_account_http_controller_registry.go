package registry

import (
	http "github.com/brianshepanek/turnbull/turnbull/output/interface/controller/enhanced_account/http"
	interactor "github.com/brianshepanek/turnbull/turnbull/output/usecase/interactor"
)

type enhancedAccountHttpControllerRegistry struct{}

func (r *registry) RegisterHttpEnhancedAccountController() {}
func (r *registry) NewHttpEnhancedAccountController() http.HttpEnhancedAccountController {
	return http.New(interactor.NewEnhancedAccountInteractor(r.newMongoEnhancedAccountRepository(), r.newDefaultEnhancedAccountPresenter()))
}
