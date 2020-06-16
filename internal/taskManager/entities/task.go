package entities

const (
	new int = iota
	completed
)

type Task struct {
	id          uint32
	name        string
	description string
	status      string
	comments    []*Comment
}

func (b *Task) Create(data map[string]string) *Task  {

}

func (b Task) ReadById(id uint32) *Task {

}

func (b *Task) UpdateById(id uint32, data map[string]string) *Task {

}

func (b *Task) DeleteById(id uint32) *Task {

}
