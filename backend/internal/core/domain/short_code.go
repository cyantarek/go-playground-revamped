package domain

type ShortCode struct {
	code string
}

func (s ShortCode) String() string {
	return s.code
}

