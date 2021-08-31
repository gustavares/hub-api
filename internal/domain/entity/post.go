package entity

type ContentType string

type Content string

type Post struct {
	id          string
	contentType ContentType
	content     Content
}

func NewPost(id string, contentType ContentType, content Content) *Post {
	return &Post{
		id,
		contentType,
		content,
	}
}
