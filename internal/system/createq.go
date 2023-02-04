package system

import "github.com/google/uuid"

func (e *Enqueted) Createq(id uuid.UUID) {
	e.ident = append(e.ident, identify{ID: id, nextID: 1})
}
