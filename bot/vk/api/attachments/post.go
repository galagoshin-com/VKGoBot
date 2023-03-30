package attachments

import "fmt"

type Post struct {
	OwnerId   int
	Id        uint
	AccessKey string
}

func (post *Post) GetType() AttachmentType {
	return PostType
}

func (post *Post) GetOwnerId() int {
	return post.OwnerId
}

func (post *Post) GetId() uint {
	return post.Id
}

func (post *Post) GetAccessKey() string {
	return post.AccessKey
}

func (post *Post) BuildString() string {
	if post.AccessKey == "" {
		return fmt.Sprintf("%s%d_%d", post.GetType(), post.OwnerId, post.Id)
	} else {
		return fmt.Sprintf("%s%d_%d_%s", post.GetType(), post.OwnerId, post.Id, post.AccessKey)
	}
}
