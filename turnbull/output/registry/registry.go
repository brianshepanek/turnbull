package registry

import http "github.com/brianshepanek/turnbull/turnbull/output/interface/controller/app/http"

type registry struct {
	accountMongoRepositoryRegistry
	accountDefaultPresenterRegistry
	enhancedAccountMongoRepositoryRegistry
	enhancedAccountDefaultPresenterRegistry
}

func New() *registry {
	return &registry{}
}
func (r *registry) NewHttpAppController() http.HttpAppController {
	r.NewHttpAccountController()
	r.NewHttpEnhancedAccountController()
	return r
}
