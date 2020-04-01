package entity

type post struct {
	postStruct
}

type posts struct {
	postsStruct
}

type Post interface {
	postInterface
}
type Posts interface {
	postsInterface
}

func NewPost() Post {
	return &post{}
}

func NewPosts() Posts {
	return &posts{}
}
