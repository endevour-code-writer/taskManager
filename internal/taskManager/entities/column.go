package entities

type Column struct {
	id     uint32
	name   string
	tasks  []*Task
}

func (b *Column) Create(data map[string]string) *Column  {

}

func (b Column) ReadById(id uint32) *Column {

}

func (b *Column) UpdateById(id uint32, data map[string]string) *Column {

}

func (b *Column) DeleteById(id uint32) *Column {

}