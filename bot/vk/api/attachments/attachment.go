package attachments

type AttachmentType string

type Attachment interface {
	GetType() AttachmentType
	GetId() uint
	GetOwnerId() int
	GetAccessKey() string
	BuildString() string
}

const (
	PhotoType = AttachmentType("photo")
	PostType  = AttachmentType("post")
)
