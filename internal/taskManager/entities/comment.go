package entities

type Comment struct {
	id   uint32
	text string
}

func (b *Comment) Create(data map[string]string) *Comment  {

}

func (b Comment) ReadById(id uint32) *Comment {

}

func (b *Comment) UpdateById(id uint32, data map[string]string) *Comment {

}

func (b *Comment) DeleteById(id uint32) *Comment {

}
