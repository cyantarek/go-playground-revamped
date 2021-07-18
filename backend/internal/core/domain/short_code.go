package domain

type ShortCode struct {
	code string
}

func NewShortCode(code string) ShortCode {
	return ShortCode{code: code}
}

func (s ShortCode) String() string {
	return s.code
}

