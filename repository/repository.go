package repository

type Repository interface {
	validate(e *Entity) error
	create(e *Entity) (*Entity, error)
	find(id int) *Entity
}
