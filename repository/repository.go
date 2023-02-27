package repository

import (
	"book-api/entity"
)

type EntityManager interface {
	Save(e *entity.Entity) (*entity.Entity, error)
	Find(id int) *entity.Entity
}

type Repository struct {
	db any
}
