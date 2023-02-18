package attachments

import "fmt"

type Image struct {
	OwnerId   int
	Id        uint
	AccessKey string
}

func (image *Image) GetType() AttachmentType {
	return Photo
}

func (image *Image) GetOwnerId() int {
	return image.OwnerId
}

func (image *Image) GetId() uint {
	return image.Id
}

func (image *Image) GetAccessKey() string {
	return image.AccessKey
}

func (image *Image) BuildString() string {
	if image.AccessKey == "" {
		return fmt.Sprintf("%s%d_%d", image.GetType(), image.OwnerId, image.Id)
	} else {
		return fmt.Sprintf("%s%d_%d_%s", image.GetType(), image.OwnerId, image.Id, image.AccessKey)
	}
}
