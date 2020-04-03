package entity

type Post struct {
	postStruct
}

func NewPost() *Post {
	return &Post{}
}

func NewPosts() *[]Post {
	return &[]Post{}
}
