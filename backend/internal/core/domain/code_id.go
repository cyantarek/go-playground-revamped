package domain

type CodeID struct {
	id string
}

func (c CodeID) String() string {
	return c.id
}

func NewCodeID(id string) CodeID {
	return CodeID{id: id}
}
