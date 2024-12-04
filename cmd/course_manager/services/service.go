package services

type Service[T any] interface {
	Create(*T) error
	All() ([]*T, error)
	ReadFromCSV() ([]*T, error)
	Update(*T) error
	Delete(*T) error
}
