package controller

import http "github.com/brianshepanek/turnbull/turnbull/output/interface/controller/account/http"

type HttpAppController interface {
	NewHttpAccountController() http.HttpAccountController
}
