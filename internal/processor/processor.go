package processor

import (
	"payment-switch-simulator/internal/model"
	"time"
)

func Process(msg model.Message) model.Response {
	// simulate processing delay
	time.Sleep(50 * time.Millisecond)

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
	}
}
