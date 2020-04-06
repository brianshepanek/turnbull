package repository

import (
	"context"
	"encoding/json"
	entity "github.com/brianshepanek/turnbull/turnbull/output/domain/entity"
	"github.com/go-redis/redis/v7"
	cast "github.com/spf13/cast"
)

type redisPostRepositoryStruct struct {
	client    *redis.Client
	namespace string
}

func (r *redisPostRepositoryStruct) Browse(ctx context.Context, req *[]entity.Post) error {

	match := r.namespace + ":*"
	keys, _, err := r.client.Scan(0, match, -1).Result()
	if err != nil {
		return err
	}

	for _, key := range keys {

		elem := entity.NewPost()
		val, err := r.client.Get(key).Result()
		if err != nil {
			return err
		}

		err = json.Unmarshal([]byte(val), elem)
		if err != nil {
			return err
		}

		*req = append(*req, *elem)

	}

	return nil

}

func (r *redisPostRepositoryStruct) Read(ctx context.Context, id int64, req *entity.Post) error {

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

func (r *redisPostRepositoryStruct) Edit(ctx context.Context, id int64, req *entity.Post) error {

	current := entity.NewPost()
	key := r.namespace + ":" + cast.ToString(id)
	currentVal, err := r.client.Get(key).Result()
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(currentVal), current)
	if err != nil {
		return err
	}

	if req.Id != nil {
		current.Id = req.Id
	}

	if req.UserId != nil {
		current.UserId = req.UserId
	}

	if req.Title != nil {
		current.Title = req.Title
	}

	if req.Subtitle != nil {
		current.Subtitle = req.Subtitle
	}

	if req.Views != nil {
		current.Views = req.Views
	}

	if req.Tags != nil {
		current.Tags = req.Tags
	}

	if req.Created != nil {
		current.Created = req.Created
	}

	if req.Modified != nil {
		current.Modified = req.Modified
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

func (r *redisPostRepositoryStruct) Add(ctx context.Context, req *entity.Post) error {

	key := r.namespace + ":" + cast.ToString(req.Id)
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

func (r *redisPostRepositoryStruct) Delete(ctx context.Context, id int64, req *entity.Post) error {

	key := r.namespace + ":" + cast.ToString(id)
	err := r.client.Del(key).Err()
	if err != nil {
		return err
	}

	return nil

}
