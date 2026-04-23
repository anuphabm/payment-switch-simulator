package async

import (
	"log"
	"time"
)

type Worker struct {
	queue *Queue
}

func NewWorker(q *Queue) *Worker {
	return &Worker{queue: q}
}

func (w *Worker) Start() {
	go func() {
		for evt := range w.queue.Subscribe() {
			processEvent(evt)
		}
	}()
}

func processEvent(evt Event) {
	start := time.Now()

	// simulate async work (e.g. callback / logging / fraud check)
	time.Sleep(20 * time.Millisecond)

	log.Printf("[ASYNC] processed event type=%s trace_id=%s latency=%s\n",
		evt.Type,
		evt.TraceID,
		time.Since(start),
	)

	maxRetry := 3

	for i := 0; i < maxRetry; i++ {
		err := handle(evt)
		if err == nil {
			return
		}
		time.Sleep(10 * time.Millisecond)
	}

	// Dead Letter Queue (mock)
	log.Printf("[DLQ] failed event trace_id=%s\n", evt.TraceID)

}

func handle(evt Event) error {
	// simulate success always (เปลี่ยนทีหลังได้)
	return nil
}
