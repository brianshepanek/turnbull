package controller

import (
	"context"
	"encoding/json"
	entity "github.com/brianshepanek/turnbull/turnbull/output/domain/entity"
	interactor "github.com/brianshepanek/turnbull/turnbull/output/usecase/interactor"
	mux "github.com/gorilla/mux"
	"net/http"
)

type httpEnhancedAccountControllerStruct struct {
	interactor interactor.EnhancedAccountInteractor
}
type httpEnhancedAccountControllerInterface interface {
	Browse(w http.ResponseWriter, r *http.Request)
	Read(w http.ResponseWriter, r *http.Request)
	Edit(w http.ResponseWriter, r *http.Request)
	Add(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func (c *httpEnhancedAccountControllerStruct) Browse(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()
	req := entity.NewEnhancedAccounts()

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

func (c *httpEnhancedAccountControllerStruct) Read(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()
	req := entity.NewEnhancedAccount()

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

func (c *httpEnhancedAccountControllerStruct) Edit(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()
	req := entity.NewEnhancedAccount()

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

func (c *httpEnhancedAccountControllerStruct) Add(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()
	req := entity.NewEnhancedAccount()

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

func (c *httpEnhancedAccountControllerStruct) Delete(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()
	req := entity.NewEnhancedAccount()

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
