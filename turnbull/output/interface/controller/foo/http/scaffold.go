package controller

import (
	"context"
	"encoding/json"
	entity "github.com/brianshepanek/turnbull/turnbull/output/domain/entity"
	interactor "github.com/brianshepanek/turnbull/turnbull/output/usecase/interactor"
	mux "github.com/gorilla/mux"
	"net/http"
	"time"
)

type httpFooControllerStruct struct {
	interactor interactor.FooInteractor
}
type httpFooControllerInterface interface {
	Browse(w http.ResponseWriter, r *http.Request)
	Read(w http.ResponseWriter, r *http.Request)
	Edit(w http.ResponseWriter, r *http.Request)
	Add(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}
type foo struct {
	entity.Foo
}

func (m *foo) MarshalJSON() ([]byte, error) {
	type jsonStructPrivate struct {
		Id       int64     `json:"id,omitempty"`
		Title    string    `json:"title,omitempty"`
		Subtitle string    `json:"subtitle,omitempty"`
		Views    int       `json:"views,omitempty"`
		Tags     []string  `json:"tags,omitempty"`
		Created  time.Time `json:"created,omitempty"`
		Modified time.Time `json:"modified,omitempty"`
	}
	jsonStruct := jsonStructPrivate{
		Created:  m.Created(),
		Id:       m.Id(),
		Modified: m.Modified(),
		Subtitle: m.Subtitle(),
		Tags:     m.Tags(),
		Title:    m.Title(),
		Views:    m.Views(),
	}
	return json.Marshal(&jsonStruct)
}

func (m *foo) UnmarshalJSON(data []byte) error {
	type jsonStructPrivate struct {
		Id       int64     `json:"id,omitempty"`
		Title    string    `json:"title,omitempty"`
		Subtitle string    `json:"subtitle,omitempty"`
		Views    int       `json:"views,omitempty"`
		Tags     []string  `json:"tags,omitempty"`
		Created  time.Time `json:"created,omitempty"`
		Modified time.Time `json:"modified,omitempty"`
	}
	jsonStruct := jsonStructPrivate{}
	err := json.Unmarshal(data, &jsonStruct)
	if err != nil {
		return err
	}
	m.SetId(jsonStruct.Id)
	m.SetTitle(jsonStruct.Title)
	m.SetSubtitle(jsonStruct.Subtitle)
	m.SetViews(jsonStruct.Views)
	m.SetTags(jsonStruct.Tags)
	m.SetCreated(jsonStruct.Created)
	m.SetModified(jsonStruct.Modified)
	return nil
}

func (c *httpFooControllerStruct) Browse(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()
	req := entity.NewFoos()

	resp, err := c.interactor.Browse(ctx, req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	var foos []*foo
	for _, elem := range resp.Elements() {
		foos = append(foos, &foo{elem})
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(foos)

}

func (c *httpFooControllerStruct) Read(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()
	foo := &foo{entity.NewFoo()}

	var stringId string
	vars := mux.Vars(r)
	if val, ok := vars["id"]; ok {
		stringId = val
	}

	id, err := foo.ToPrimary(ctx, stringId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	resp, err := c.interactor.Read(ctx, id, foo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (c *httpFooControllerStruct) Edit(w http.ResponseWriter, r *http.Request) {}

func (c *httpFooControllerStruct) Add(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()
	foo := &foo{entity.NewFoo()}

	err := json.NewDecoder(r.Body).Decode(foo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	resp, err := c.interactor.Add(ctx, foo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)

}

func (c *httpFooControllerStruct) Delete(w http.ResponseWriter, r *http.Request) {}
