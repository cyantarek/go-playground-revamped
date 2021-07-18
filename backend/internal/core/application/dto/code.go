package dto

import (
	"time"
)

type FormatCode struct {
	Code     string
	Language string
}

type FormatCodeResult struct {
	Code string
}

type RunCode struct {
	Code     string
	Language string
}

type RunCodeResult struct {
	RunID   string
	Output  string
	RunTime float64
}

type ShareCode struct {
	Code     string
	Language string
}

type ShareCodeResult struct {
	ShortCode string
}

type GetCodeByID struct {
	ID        string
	Code      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
