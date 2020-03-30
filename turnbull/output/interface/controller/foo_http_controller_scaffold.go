package controller

import (
	"context"
	"encoding/json"
	entity "github.com/brianshepanek/turnbull/turnbull/output/domain/entity"
	interactor "github.com/brianshepanek/turnbull/turnbull/output/usecase/interactor"
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
		Id       string    `json:"id,omitempty"`
		Title    string    `json:"title,omitempty"`
		Subtitle string    `json:"subtitle,omitempty"`
		Int      int       `json:"int,omitempty"`
		Tags     []string  `json:"tags,omitempty"`
		Created  time.Time `json:"created,omitempty"`
	}
	jsonStruct := jsonStructPrivate{
		Created:  m.Created(),
		Id:       m.Id(),
		Int:      m.Int(),
		Subtitle: m.Subtitle(),
		Tags:     m.Tags(),
		Title:    m.Title(),
	}
	return json.Marshal(&jsonStruct)
}

func (m *foo) UnmarshalJSON(data []byte) error {
	type jsonStructPrivate struct {
		Id       string    `json:"id,omitempty"`
		Title    string    `json:"title,omitempty"`
		Subtitle string    `json:"subtitle,omitempty"`
		Int      int       `json:"int,omitempty"`
		Tags     []string  `json:"tags,omitempty"`
		Created  time.Time `json:"created,omitempty"`
	}
	jsonStruct := jsonStructPrivate{}
	err := json.Unmarshal(data, &jsonStruct)
	if err != nil {
		return err
	}
	m.SetId(jsonStruct.Id)
	m.SetTitle(jsonStruct.Title)
	m.SetSubtitle(jsonStruct.Subtitle)
	m.SetInt(jsonStruct.Int)
	m.SetTags(jsonStruct.Tags)
	m.SetCreated(jsonStruct.Created)
	return nil
}

func (c *httpFooControllerStruct) Browse(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	req := entity.NewFoos()
	resp, err := c.interactor.Browse(ctx, nil, req)
	if err != nil {
	}
	var foos []*foo
	for _, elem := range resp.Elements() {
		foos = append(foos, &foo{elem})
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp.Elements())
}
func (c *httpFooControllerStruct) Read(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	req := entity.NewFoo()
	resp, err := c.interactor.Read(ctx, nil, req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
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
	}
	resp, err := c.interactor.Add(ctx, foo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
func (c *httpFooControllerStruct) Delete(w http.ResponseWriter, r *http.Request) {}
