package controller

import (
	"encoding/json"
	interactor "github.com/brianshepanek/turnbull/turnbull/output/scaffold/usecase/interactor"
	"net/http"
)

type httpFooScaffoldController struct {
	interactor interactor.FooScaffoldInteractor
}

func NewHttpFooScaffoldController(interactor interactor.FooScaffoldInteractor) *httpFooScaffoldController {
	return &httpFooScaffoldController{interactor}
}
func (c *httpFooScaffoldController) Browse(w http.ResponseWriter, r http.Request) {
	var resp interface{}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
func (c *httpFooScaffoldController) Read(w http.ResponseWriter, r http.Request) {
	var resp interface{}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
func (c *httpFooScaffoldController) Edit(w http.ResponseWriter, r http.Request) {
	var resp interface{}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
func (c *httpFooScaffoldController) Add(w http.ResponseWriter, r http.Request) {
	var resp interface{}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
func (c *httpFooScaffoldController) Delete(w http.ResponseWriter, r http.Request) {
	var resp interface{}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
