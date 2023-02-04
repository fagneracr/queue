package system

func (e *Enqueted) Createq(id string) {
	e.ident = append(e.ident, Identify{ID: id, NextID: 1})
}
