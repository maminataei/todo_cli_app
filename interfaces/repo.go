package interfaces

type Repo[T any] interface {
	Create(T) error
	ListAll() ([]T, error)
	Get(id int) (T, error)
	Edit(T) error
	Delete(id int) error
}
