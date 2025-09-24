package repo

type CrudEntity[T any] interface {
	Create(entity T) error
	GetById(id string) (*T, error)
	Delete(id string) error
	Update(entity T) error
	GetAll() ([]T, error)
}
