package registry

import http "github.com/brianshepanek/turnbull/turnbull/output/interface/controller/app/http"

type registry struct {
	accountMongoRepositoryRegistry
	accountDefaultPresenterRegistry
}

func New() *registry {
	return &registry{}
}
func (r *registry) NewHttpAppController() http.HttpAppController {
	r.NewHttpAccountController()
	return r
}
