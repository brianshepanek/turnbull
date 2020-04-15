package controller

import (
	http "github.com/brianshepanek/turnbull/turnbull/output/interface/controller/account/http"
	http2 "github.com/brianshepanek/turnbull/turnbull/output/interface/controller/channel/http"
	http3 "github.com/brianshepanek/turnbull/turnbull/output/interface/controller/message/http"
	http1 "github.com/brianshepanek/turnbull/turnbull/output/interface/controller/user/http"
)

type HttpAppController interface {
	NewHttpAccountController() http.HttpAccountController
	NewHttpUserController() http1.HttpUserController
	NewHttpChannelController() http2.HttpChannelController
	NewHttpMessageController() http3.HttpMessageController
}
