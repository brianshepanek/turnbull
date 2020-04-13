package registry

import http "github.com/brianshepanek/turnbull/turnbull/output/interface/controller/user/http"

type userHttpControllerRegistry struct{}

func (r *registry) RegisterHttpUserController() {}
func (r *registry) NewHttpUserController() http.HttpUserController {
	return http.New(r.newUserInteractor())
}
