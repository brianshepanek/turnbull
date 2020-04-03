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

type httpCommentControllerStruct struct {
	interactor interactor.CommentInteractor
}
type httpCommentControllerInterface interface {
	Browse(w http.ResponseWriter, r *http.Request)
	Read(w http.ResponseWriter, r *http.Request)
	Edit(w http.ResponseWriter, r *http.Request)
	Add(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}
type commentLocal struct {
	entity.Comment
}

func (m *commentLocal) MarshalJSON() ([]byte, error) {
	type jsonStructPrivate struct {
		Id       *int64     `json:"id,omitempty"`
		PostId   *int64     `json:"post_id,omitempty"`
		Title    *string    `json:"title,omitempty"`
		Body     *string    `json:"body,omitempty"`
		Created  *time.Time `json:"created,omitempty"`
		Modified *time.Time `json:"modified,omitempty"`
	}
	jsonStruct := jsonStructPrivate{
		Body:     m.Body(),
		Created:  m.Created(),
		Id:       m.Id(),
		Modified: m.Modified(),
		PostId:   m.PostId(),
		Title:    m.Title(),
	}
	return json.Marshal(&jsonStruct)
}

func (m *commentLocal) UnmarshalJSON(data []byte) error {
	type jsonStructPrivate struct {
		Id       *int64     `json:"id,omitempty"`
		PostId   *int64     `json:"post_id,omitempty"`
		Title    *string    `json:"title,omitempty"`
		Body     *string    `json:"body,omitempty"`
		Created  *time.Time `json:"created,omitempty"`
		Modified *time.Time `json:"modified,omitempty"`
	}
	jsonStruct := jsonStructPrivate{}
	err := json.Unmarshal(data, &jsonStruct)
	if err != nil {
		return err
	}
	m.SetId(jsonStruct.Id)
	m.SetPostId(jsonStruct.PostId)
	m.SetTitle(jsonStruct.Title)
	m.SetBody(jsonStruct.Body)
	m.SetCreated(jsonStruct.Created)
	m.SetModified(jsonStruct.Modified)
	return nil
}

func (c *httpCommentControllerStruct) Browse(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()
	req := entity.NewComments()

	resp, err := c.interactor.Browse(ctx, req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	var comments []*commentLocal
	for _, elem := range resp.Elements() {
		comments = append(comments, &commentLocal{elem})
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(comments)

}

func (c *httpCommentControllerStruct) Read(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()
	req := &commentLocal{entity.NewComment()}

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
	json.NewEncoder(w).Encode(&commentLocal{resp})

}

func (c *httpCommentControllerStruct) Edit(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()
	req := &commentLocal{entity.NewComment()}

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
	json.NewEncoder(w).Encode(&commentLocal{resp})

}

func (c *httpCommentControllerStruct) Add(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()
	req := &commentLocal{entity.NewComment()}

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
	json.NewEncoder(w).Encode(&commentLocal{resp})

}

func (c *httpCommentControllerStruct) Delete(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()
	req := &commentLocal{entity.NewComment()}

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
