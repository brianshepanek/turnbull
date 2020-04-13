package controller

import (
	"context"
	"encoding/json"
	entity "github.com/brianshepanek/turnbull/turnbull/output/domain/entity"
	interactor "github.com/brianshepanek/turnbull/turnbull/output/usecase/interactor"
	mux "github.com/gorilla/mux"
	"net/http"
)

type httpUserControllerStruct struct {
	interactor interactor.UserInteractor
}
type httpUserControllerInterface interface {
	Browse(w http.ResponseWriter, r *http.Request)
	BrowseByAccountId(w http.ResponseWriter, r *http.Request)
	Read(w http.ResponseWriter, r *http.Request)
	Edit(w http.ResponseWriter, r *http.Request)
	Add(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func (c *httpUserControllerStruct) Browse(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()
	req := entity.NewUsers()

	resp, err := c.interactor.Browse(ctx, req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp.Elements())

}

func (c *httpUserControllerStruct) BrowseByAccountId(w http.ResponseWriter, r *http.Request) {}

func (c *httpUserControllerStruct) Read(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()
	req := entity.NewUser()

	var stringId string
	vars := mux.Vars(r)
	if val, ok := vars["id"]; ok {
		stringId = val
	}

	id, err := req.ToPrimary(ctx, stringId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	resp, err := c.interactor.Read(ctx, id, req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&resp)

}

func (c *httpUserControllerStruct) Edit(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()
	req := entity.NewUser()

	var stringId string
	vars := mux.Vars(r)
	if val, ok := vars["id"]; ok {
		stringId = val
	}

	id, err := req.ToPrimary(ctx, stringId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	err = json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	resp, err := c.interactor.Edit(ctx, id, req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&resp)

}

func (c *httpUserControllerStruct) Add(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()
	req := entity.NewUser()

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	resp, err := c.interactor.Add(ctx, req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&resp)

}

func (c *httpUserControllerStruct) Delete(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()
	req := entity.NewUser()

	var stringId string
	vars := mux.Vars(r)
	if val, ok := vars["id"]; ok {
		stringId = val
	}

	id, err := req.ToPrimary(ctx, stringId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	_, err = c.interactor.Delete(ctx, id, req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

}
