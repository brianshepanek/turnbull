package controller

import (
	http2 "github.com/brianshepanek/turnbull/turnbull/output/interface/controller/account/http"
	http1 "github.com/brianshepanek/turnbull/turnbull/output/interface/controller/lucky/http"
	http "github.com/brianshepanek/turnbull/turnbull/output/interface/controller/model/http"
)

type HttpAppController interface {
	NewHttpModelController() http.HttpModelController
	NewHttpLuckyController() http1.HttpLuckyController
	NewHttpAccountController() http2.HttpAccountController
}
