package async

type Event struct {
	TraceID string
	Type    string
	Payload []byte
}
