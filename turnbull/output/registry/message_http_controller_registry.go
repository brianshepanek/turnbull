package registry

import http "github.com/brianshepanek/turnbull/turnbull/output/interface/controller/message/http"

type messageHttpControllerRegistry struct{}

func (r *registry) RegisterHttpMessageController() {}
func (r *registry) NewHttpMessageController() http.HttpMessageController {
	return http.New(r.newMessageInteractor())
}
