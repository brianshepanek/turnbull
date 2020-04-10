package registry

import (
	http "github.com/brianshepanek/turnbull/turnbull/output/interface/controller/account/http"
	interactor "github.com/brianshepanek/turnbull/turnbull/output/usecase/interactor"
)

type accountHttpControllerRegistry struct{}

func (r *registry) RegisterHttpAccountController() {}
func (r *registry) NewHttpAccountController() http.HttpAccountController {
	return http.New(interactor.NewAccountInteractor(r.newMongoAccountRepository(), r.newDefaultAccountPresenter()))
}
