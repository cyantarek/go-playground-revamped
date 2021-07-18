package domain

import (
	"time"
)

type Code struct {
	id        CodeID
	code      string
	visited   int
	createdAt time.Time
	updatedAt time.Time
}

func (c *Code) MarkVisit() {
	c.visited++
}

func (c Code) VisitedHowManyTimes() int {
	return c.visited
}

func (c Code) CodeID() CodeID {
	return c.id
}

func NewCode(id CodeID, body string) Code {
	return Code{id: id, code: body, createdAt: time.Now(), updatedAt: time.Now()}
}

func (c *Code) TimeWise(createdAt, updatedAt time.Time) {
	// time validation

	// time related business logic

	c.createdAt = createdAt
	c.updatedAt = updatedAt
}

func (c Code) Code() string {
	return c.code
}

func (c Code) WhenCreated() time.Time {
	return c.createdAt
}

func (c Code) WhenLastUpdated() time.Time {
	return c.updatedAt
}
