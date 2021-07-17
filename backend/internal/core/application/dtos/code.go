package dtos

type FormatCode struct {
	Body     string
	Language string
}

type FormatCodeResult struct {
	Body string
}

type RunCode struct {
	Body     string
	Language string
}

type RunCodeResult struct {
	RunID   string
	Status  string
	Output  string
	Error   string
	RunTime float64
}
