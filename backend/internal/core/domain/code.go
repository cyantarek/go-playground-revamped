package domain

import (
	"time"
)

type Code struct {
	id        CodeID
	body      string
	shortCode ShortCode
	createdAt time.Time
	updatedAt time.Time
}

func (c Code) Id() CodeID {
	return c.id
}

func NewCode(body string) Code {
	return Code{body: body, createdAt: time.Now(), updatedAt: time.Now()}
}

func (c *Code) Shorten(shortCode ShortCode) {
	c.shortCode = shortCode
}

func (c Code) Body() string {
	return c.body
}

func (c Code) ShortCode() ShortCode {
	return c.shortCode
}

func (c Code) CreatedAt() time.Time {
	return c.createdAt
}

func (c Code) UpdatedAt() time.Time {
	return c.updatedAt
}
