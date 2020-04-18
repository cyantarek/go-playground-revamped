package playground

// Service API Requests Responses
type PingResponse struct {
	Message string `json:"message"`
}

type CodeByShareRequest struct {
	Id string `json:"id"`
}

type CodeByShareResponse struct {
	Body      string `json:"body"`
	ShareCode string `json:"share_code"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type CommonCodeRequest struct {
	Body string `json:"body"`
}

type ShaCodeResponse struct {
	Code string `json:"code"`
}

type FormatCodeResponse struct {
	FormattedCode string `json:"formatted_code"`
}

type CodeRunResponse struct {
	Status  string  `json:"status"`
	Output  string  `json:"output"`
	Error   string  `json:"error"`
	RunTime float64 `json:"run_time"`
}
