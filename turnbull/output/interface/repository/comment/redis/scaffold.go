package repository

import (
	"context"
	"encoding/json"
	entity "github.com/brianshepanek/turnbull/turnbull/output/domain/entity"
	"github.com/go-redis/redis/v7"
	cast "github.com/spf13/cast"
)

type redisCommentRepositoryStruct struct {
	client    *redis.Client
	namespace string
}

func (r *redisCommentRepositoryStruct) Browse(ctx context.Context, req entity.Comments) error {

	match := r.namespace + ":*"
	keys, _, err := r.client.Scan(0, match, -1).Result()
	if err != nil {
		return err
	}

	for _, key := range keys {

		elem := entity.NewComment()
		val, err := r.client.Get(key).Result()
		if err != nil {
			return err
		}

		err = json.Unmarshal([]byte(val), elem)
		if err != nil {
			return err
		}

		req.Append(elem)

	}

	return nil

}

func (r *redisCommentRepositoryStruct) Read(ctx context.Context, id int64, req entity.Comment) error {

	key := r.namespace + ":" + cast.ToString(id)
	val, err := r.client.Get(key).Result()
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(val), req)
	if err != nil {
		return err
	}

	return nil

}

func (r *redisCommentRepositoryStruct) Edit(ctx context.Context, id int64, req entity.Comment) error {

	current := entity.NewComment()
	key := r.namespace + ":" + cast.ToString(id)
	currentVal, err := r.client.Get(key).Result()
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(currentVal), current)
	if err != nil {
		return err
	}

	if req.Id() != nil {
		current.SetId(req.Id())
	}

	if req.PostId() != nil {
		current.SetPostId(req.PostId())
	}

	if req.UserId() != nil {
		current.SetUserId(req.UserId())
	}

	if req.Title() != nil {
		current.SetTitle(req.Title())
	}

	if req.Body() != nil {
		current.SetBody(req.Body())
	}

	if req.Created() != nil {
		current.SetCreated(req.Created())
	}

	if req.Modified() != nil {
		current.SetModified(req.Modified())
	}

	val, err := json.Marshal(current)
	if err != nil {
		return err
	}

	err = r.client.Set(key, val, 0).Err()
	if err != nil {
		return err
	}

	return nil

}

func (r *redisCommentRepositoryStruct) Add(ctx context.Context, req entity.Comment) error {

	key := r.namespace + ":" + cast.ToString(req.Id())
	val, err := json.Marshal(req)
	if err != nil {
		return err
	}

	err = r.client.Set(key, val, 0).Err()
	if err != nil {
		return err
	}

	return nil

}

func (r *redisCommentRepositoryStruct) Delete(ctx context.Context, id int64, req entity.Comment) error {

	key := r.namespace + ":" + cast.ToString(id)
	err := r.client.Del(key).Err()
	if err != nil {
		return err
	}

	return nil

}
