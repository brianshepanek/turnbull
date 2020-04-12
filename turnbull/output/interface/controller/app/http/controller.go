package controller

import (
	http "github.com/brianshepanek/turnbull/turnbull/output/interface/controller/account/http"
	http1 "github.com/brianshepanek/turnbull/turnbull/output/interface/controller/enhanced_account/http"
)

type HttpAppController interface {
	NewHttpAccountController() http.HttpAccountController
	NewHttpEnhancedAccountController() http1.HttpEnhancedAccountController
}
