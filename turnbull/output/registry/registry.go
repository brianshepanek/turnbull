package registry

import http "github.com/brianshepanek/turnbull/turnbull/output/interface/controller/app/http"

type registry struct {
	modelMongoRepositoryRegistry
	modelDefaultPresenterRegistry
	accountMongoRepositoryRegistry
	accountDefaultPresenterRegistry
}

func New() *registry {
	return &registry{}
}
func (r *registry) NewHttpAppController() http.HttpAppController {
	r.NewHttpModelController()
	r.NewHttpAccountController()
	return r
}
