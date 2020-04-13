package registry

import http "github.com/brianshepanek/turnbull/turnbull/output/interface/controller/account/http"

type accountHttpControllerRegistry struct{}

func (r *registry) RegisterHttpAccountController() {}
func (r *registry) NewHttpAccountController() http.HttpAccountController {
	return http.New(r.newAccountInteractor())
}
