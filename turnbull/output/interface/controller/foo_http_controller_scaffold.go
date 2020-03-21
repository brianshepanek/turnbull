package controller

import (
	"context"
	"encoding/json"
	entity "github.com/brianshepanek/turnbull/turnbull/output/domain/entity"
	interactor "github.com/brianshepanek/turnbull/turnbull/output/usecase/interactor"
	"net/http"
)

type httpFooControllerStruct struct {
	interactor interactor.FooInteractor
}
type httpFooControllerInterface interface {
	Browse(w http.ResponseWriter, r http.Request)
	Read(w http.ResponseWriter, r http.Request)
	Edit(w http.ResponseWriter, r http.Request)
	Add(w http.ResponseWriter, r http.Request)
	Delete(w http.ResponseWriter, r http.Request)
}

func (c *httpFooControllerStruct) Browse(w http.ResponseWriter, r http.Request) {
	ctx := context.Background()
	req := entity.NewFoos()
	resp, err := c.interactor.Browse(ctx, nil, req)
	if err != nil {
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
func (c *httpFooControllerStruct) Read(w http.ResponseWriter, r http.Request) {
	ctx := context.Background()
	req := entity.NewFoo()
	resp, err := c.interactor.Read(ctx, nil, req)
	if err != nil {
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
func (c *httpFooControllerStruct) Edit(w http.ResponseWriter, r http.Request) {}
func (c *httpFooControllerStruct) Add(w http.ResponseWriter, r http.Request) {
	ctx := context.Background()
	req := entity.NewFoo()
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
	}
	resp, err := c.interactor.Add(ctx, req)
	if err != nil {
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
func (c *httpFooControllerStruct) Delete(w http.ResponseWriter, r http.Request) {}
