package handler

import (
	"encoding/json"
	"log"
	"payment-switch-simulator/internal/async"
	"payment-switch-simulator/internal/model"
	"payment-switch-simulator/internal/processor"
	"payment-switch-simulator/internal/store"
	"time"
)

var memStore = store.NewMemoryStore()

func HandleMessage(raw string, q *async.Queue) string {
	start := time.Now()

	var msg model.Message
	err := json.Unmarshal([]byte(raw), &msg)
	if err != nil {
		return `{"status":"error","message":"invalid format"}`
	}

	// ✅ Idempotency check
	if rec, found := memStore.Get(msg.TraceID); found {
		log.Printf("DUPLICATE TX=%s -> returning cached response\n", msg.TraceID)
		return string(rec.Response)
	}

	// process normally
	resp := processor.Process(msg)
	resBytes, _ := json.Marshal(resp)

	// ✅ Publish event for async processing (e.g. logging, metrics)
	evt := async.Event{
		TraceID: msg.TraceID,
		Type:    "TX_PROCESSED",
		Payload: resBytes,
	}

	q.Publish(evt)

	// ✅ Save result for future duplicate requests
	memStore.Set(msg.TraceID, resBytes)

	latency := time.Since(start)
	log.Printf("TX=%s MTI=%s LATENCY=%s\n", msg.TraceID, msg.MTI, latency)

	return string(resBytes)
}
