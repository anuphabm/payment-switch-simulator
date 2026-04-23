package handler

import (
	"encoding/json"
	"log"
	"payment-switch-simulator/internal/model"
	"payment-switch-simulator/internal/processor"
	"time"
)

func HandleMessage(raw string) string {
	start := time.Now()

	var msg model.Message
	err := json.Unmarshal([]byte(raw), &msg)
	if err != nil {
		return `{"status":"error","message":"invalid format"}`
	}

	resp := processor.Process(msg)

	latency := time.Since(start)
	log.Printf("TX=%s MTI=%s LATENCY=%s\n", msg.TraceID, msg.MTI, latency)

	resBytes, _ := json.Marshal(resp)
	return string(resBytes)
}
