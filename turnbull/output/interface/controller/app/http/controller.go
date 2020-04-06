package controller

import (
	http1 "github.com/brianshepanek/turnbull/turnbull/output/interface/controller/comment/http"
	http "github.com/brianshepanek/turnbull/turnbull/output/interface/controller/post/http"
)

type HttpAppController interface {
	NewHttpPostController() http.HttpPostController
	NewHttpCommentController() http1.HttpCommentController
}
