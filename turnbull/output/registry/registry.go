package registry

import http "github.com/brianshepanek/turnbull/turnbull/output/interface/controller/app/http"

type registry struct {
	accountMongoRepositoryRegistry
	accountDefaultPresenterRegistry
	userMongoRepositoryRegistry
	userDefaultPresenterRegistry
}

func New() *registry {
	return &registry{}
}
func (r *registry) NewHttpAppController() http.HttpAppController {
	r.NewHttpAccountController()
	r.NewHttpUserController()
	return r
}
