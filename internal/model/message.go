package model

type Message struct {
	MTI     string `json:"mti"`
	TraceID string `json:"trace_id"`
	Amount  int    `json:"amount"`
}

type Response struct {
	MTI      string `json:"mti"`
	TraceID  string `json:"trace_id"`
	Response string `json:"response_code"`
	Status   string `json:"status,omitempty"`
}
