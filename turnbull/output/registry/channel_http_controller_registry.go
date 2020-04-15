package registry

import http "github.com/brianshepanek/turnbull/turnbull/output/interface/controller/channel/http"

type channelHttpControllerRegistry struct{}

func (r *registry) RegisterHttpChannelController() {}
func (r *registry) NewHttpChannelController() http.HttpChannelController {
	return http.New(r.newChannelInteractor())
}
