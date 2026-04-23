package processor

import (
	"errors"
	"math/rand"
	"payment-switch-simulator/internal/model"
	"payment-switch-simulator/internal/resilience"
	"time"
)

var cb = resilience.NewCircuitBreaker(3, 5*time.Second)

func Process(msg model.Message) model.Response {

	err := cb.Execute(func() error {
		// simulate downstream call (เช่น core banking)
		if rand.Intn(10) < 3 { // 30% fail
			return errors.New("downstream error")
		}
		time.Sleep(50 * time.Millisecond)
		return nil
	})

	if err != nil {
		return model.Response{
			MTI:      "0210",
			TraceID:  msg.TraceID,
			Response: "91", // issuer unavailable
			Status:   "downstream error",
		}
	}

	if msg.MTI == "0200" {
		return model.Response{
			MTI:      "0210",
			TraceID:  msg.TraceID,
			Response: "00", // approved
		}
	}

	return model.Response{
		MTI:      "9999",
		TraceID:  msg.TraceID,
		Response: "96", // system error
		Status:   "system error",
	}
}
