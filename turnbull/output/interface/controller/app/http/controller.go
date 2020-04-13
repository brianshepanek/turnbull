package controller

import (
	http "github.com/brianshepanek/turnbull/turnbull/output/interface/controller/account/http"
	http1 "github.com/brianshepanek/turnbull/turnbull/output/interface/controller/user/http"
)

type HttpAppController interface {
	NewHttpAccountController() http.HttpAccountController
	NewHttpUserController() http1.HttpUserController
}
