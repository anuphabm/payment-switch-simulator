# Payment Switch Simulator (ISO8583 Mock)

A production-inspired simulation of a payment transaction switch, designed to demonstrate how real-world financial systems (POS/EDC/ATM networks) process transactions with reliability, consistency, and low latency.

This project focuses on **system design, transaction flow integrity,** and backend architecture,rather than just code implementation.

---

## 🧠 Why This Project Exists

Most real-world payment systems operate under strict requirements:

- **No duplicate financial transactions (Idempotency)**
- **Low latency (sub-second response time)**
- **High concurrency handling**
- **Reliable request/response guarantees**
- **Strict protocol handling (ISO8583)**

This repository simulates those concerns in a simplified but architecturally accurate way.

---

## 🚀 Key Capabilities

- TCP-based transaction processing (simulating POS/EDC communication)
- ISO8583-inspired message flow (0200 → 0210)
- Idempotency control (duplicate transaction prevention)
- Concurrent connection handling
- Transaction latency tracking
- In-memory transaction store (for fast response + replay)
- Mock Kafka-style asynchronous processing (channel-based event queue)
- Timeout-ready architecture (extensible)

---

## 🔁 Transaction Flow

1. Client sends transaction request (MTI: 0200)
2. System parses incoming message
3. Idempotency check is performed using `trace_id`
4. If duplicate → return cached response
5. If new → process transaction
6. Generate response (MTI: 0210)
7. Store result for future replay
8. Publish event to async queue (Kafka-style mock)
9. Background worker processes event (logging / callback simulation)
10. Return response to client

---

## 🧠 Architecture Overview

The system follows a layered architecture commonly used in payment switching systems:

- **Transport Layer**: Handles TCP connections
  - Handles TCP connections
  - Responsible for network I/O and connection lifecycle
- **Handler Layer**: Parses and routes messages
  - Parses incoming messages
  - Applies idempotency logic
  - Routes to business processor
- **Processor Layer**: Executes business logic
  - Executes business logic (authorization simulation)
  - Generates response codes
- **Store Layer**: Simulates transaction persistence
  - Maintains transaction state (in-memory)
  - Ensures idempotent behavior

See: `/docs/architecture.md`

---

## ⚡ Async Flow (Kafka Simulation)

### Architecture

```text
Client → TCP → Handler → Processor → Response
                               ↓
                          Event Producer
                               ↓
                        Channel (Mock Kafka)
                               ↓
                          Worker Consumer
```

### Design Insight

Real systems use Kafka for:

- durability
- replay capability
- decoupled services

This project simulates that using Go channels to demonstrate:

- event-driven architecture
- non-blocking processing
- horizontal scalability concepts

---

## 🧩 Design Decisions

### 1. Idempotency (Critical for Financial Systems)

Duplicate requests (e.g., due to network retry) must not result in duplicate processing.

This is implemented using:

- `trace_id` as a unique transaction key
- In-memory store with response caching

---

### 2. Layer Separation

Each component has a single responsibility:

- Transport ≠ Business Logic
- Handler ≠ Processor
  
This allow:

- Easier scaling
- Better testability
- Future extensibility (e.g., plug Kafka, DB)
  
---

### 3. Concurrency Model

- Each connection is handled in a separate goroutine
- Shared state is protected using mutex (thread-safe store)

---

### 4. Latency Awareness

Each transaction is measured and logged:

- Helps simulate SLA monitoring
- Critical in real-world financial systems
  
---

### 📊 Example Request

`{"mti":"0200","trace_id":"123456","amount":100}`

### 📊 Example Response

`{"mti":"0210","trace_id":"123456","response_code":"00"}`

---

## 🔐 Real-World Considerations (Not Implemented Here)

This project is intentionally simplified.
Production systems would include:

- Full ISO8583 bitmap & field parsing
- HSM integration (encryption / PIN handling)
- PCI-DSS compliance
- Fraud detection systems
- Distributed transaction store (e.g., Redis, DB)
- Message queue (Kafka / MQ) for async processing
- Load balancing and failover

---

## 📈 Scaling Strategy (Production Scenario)

If scaling to production, the system would evolve as follows:

- Stateless service instances behind Load Balancer
- Kafka cluster replacing channel queue
- Idempotency store moved to Redis / distributed cache
- Kafka introduced for async processing & retry
- Circuit breaker & retry policies added
- Consumer groups for scaling
- Observability stack (Prometheus, ELK, tracing)
  
---

## ⚙️ Tech Stack

- Go (Golang)
- TCP Socket Programming
- JSON (for simplified message structure)
- Docker (optional)

---

## 🧪 Running the Project

`go run cmd/server/main.go`

Test via:

`echo '{"mti":"0200","trace_id":"ABC123","amount":100}' | nc localhost 9090`

---

## 📌 Notes

This project is for demonstration purposes only.  
Most real-world implementations involve secure ISO8583 parsing, HSM integration, and strict compliance (PCI-DSS).

---

## 👤 Author

Most of my production work is in private enterprise repositories, particularly in banking and payment systems.  
This repository is created to demonstrate architectural approach and system design thinking.
