package resilience

import (
	"errors"
	"sync"
	"time"
)

type State int

const (
	CLOSED State = iota
	OPEN
	HALF_OPEN
)

type CircuitBreaker struct {
	mu sync.Mutex

	state State

	failureCount int
	successCount int

	failureThreshold int
	recoveryTimeout  time.Duration

	lastFailureTime time.Time
}

func NewCircuitBreaker(threshold int, timeout time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		state:            CLOSED,
		failureThreshold: threshold,
		recoveryTimeout:  timeout,
	}
}

func (cb *CircuitBreaker) Execute(fn func() error) error {
	cb.mu.Lock()

	// check OPEN state
	if cb.state == OPEN {
		if time.Since(cb.lastFailureTime) > cb.recoveryTimeout {
			cb.state = HALF_OPEN
		} else {
			cb.mu.Unlock()
			return errors.New("circuit breaker is OPEN")
		}
	}

	cb.mu.Unlock()

	// execute function
	err := fn()

	cb.mu.Lock()
	defer cb.mu.Unlock()

	if err != nil {
		cb.failureCount++
		cb.lastFailureTime = time.Now()

		if cb.failureCount >= cb.failureThreshold {
			cb.state = OPEN
		}
		return err
	}

	// success case
	cb.failureCount = 0

	if cb.state == HALF_OPEN {
		cb.successCount++
		if cb.successCount >= 2 {
			cb.state = CLOSED
			cb.successCount = 0
		}
	}

	return nil
}
