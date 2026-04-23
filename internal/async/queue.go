package async

type Queue struct {
	ch chan Event
}

func NewQueue(buffer int) *Queue {
	return &Queue{
		ch: make(chan Event, buffer),
	}
}

// Producer
func (q *Queue) Publish(evt Event) {
	q.ch <- evt
}

// Consumer
func (q *Queue) Subscribe() <-chan Event {
	return q.ch
}
