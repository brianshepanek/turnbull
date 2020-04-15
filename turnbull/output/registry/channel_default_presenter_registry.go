package registry

import (
	default1 "github.com/brianshepanek/turnbull/turnbull/output/interface/presenter/channel/default"
	presenter "github.com/brianshepanek/turnbull/turnbull/output/usecase/presenter"
)

type channelDefaultPresenterRegistry struct{}

func (r *registry) RegisterDefaultChannelPresenter() {}
func (r *registry) newDefaultChannelPresenter() presenter.ChannelPresenter {
	return default1.New()
}
