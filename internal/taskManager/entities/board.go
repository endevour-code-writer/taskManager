package entities

type Board struct {
	id          uint32
	name        string
	description string
	columns     []*Column
}

func (b *Board) Create(data map[string]string) *Board  {

}

func (b Board) ReadById(id uint32) *Board {

}

func (b *Board) UpdateById(id uint32, data map[string]string) *Board {

}

func (b *Board) DeleteById(id uint32) *Board {

}
