package todo

type Repository interface {
	Create(t *Todo) error
	GetAll() ([]Todo, error)
	GetByID(id int) (*Todo, error)
}
