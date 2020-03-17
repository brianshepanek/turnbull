package controller

import (
	"context"
	"encoding/json"
	entity "github.com/brianshepanek/turnbull/turnbull/output/scaffold/domain/entity"
	interactor "github.com/brianshepanek/turnbull/turnbull/output/scaffold/usecase/interactor"
	"net/http"
)

type httpFooScaffoldController struct {
	interactor interactor.FooScaffoldInteractor
}
type HttpFooScaffoldController interface {
	Browse(w http.ResponseWriter, r http.Request)
	Read(w http.ResponseWriter, r http.Request)
	Edit(w http.ResponseWriter, r http.Request)
	Add(w http.ResponseWriter, r http.Request)
	Delete(w http.ResponseWriter, r http.Request)
}

func NewHttpFooScaffoldController(interactor interactor.FooScaffoldInteractor) *httpFooScaffoldController {
	return &httpFooScaffoldController{interactor}
}
func (c *httpFooScaffoldController) Browse(w http.ResponseWriter, r http.Request) {
	ctx := context.Background()
	req := entity.NewFoosScaffoldStruct()
	resp, err := c.interactor.Browse(ctx, nil, req)
	if err != nil {
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
func (c *httpFooScaffoldController) Read(w http.ResponseWriter, r http.Request) {
	ctx := context.Background()
	req := entity.NewFooScaffoldStruct()
	resp, err := c.interactor.Read(ctx, nil, req)
	if err != nil {
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
func (c *httpFooScaffoldController) Edit(w http.ResponseWriter, r http.Request) {}
func (c *httpFooScaffoldController) Add(w http.ResponseWriter, r http.Request) {
	ctx := context.Background()
	req := entity.NewFooScaffoldStruct()
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
	}
	resp, err := c.interactor.Add(ctx, req)
	if err != nil {
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
func (c *httpFooScaffoldController) Delete(w http.ResponseWriter, r http.Request) {}
