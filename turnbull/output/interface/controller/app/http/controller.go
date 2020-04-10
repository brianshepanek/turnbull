package controller

import (
	http1 "github.com/brianshepanek/turnbull/turnbull/output/interface/controller/account/http"
	http "github.com/brianshepanek/turnbull/turnbull/output/interface/controller/model/http"
)

type HttpAppController interface {
	NewHttpModelController() http.HttpModelController
	NewHttpAccountController() http1.HttpAccountController
}
