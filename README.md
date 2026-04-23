# Payment Switch Simulator (ISO8583 Mock)

This project is a simplified simulation of a payment transaction switch, inspired by real-world financial systems such as POS/EDC and ATM networks.

It demonstrates how transaction messages are processed over TCP using a simplified ISO8583-like protocol.

---

## 🚀 Key Features

- TCP-based transaction processing
- ISO8583-like message parsing (mock implementation)
- Request/Response flow (0200 → 0210)
- Timeout & retry handling
- Concurrent request processing
- Structured logging with latency tracking

---

## 🧠 Architecture Overview

The system is designed using a layered architecture:

- **Transport Layer**: Handles TCP connections
- **Handler Layer**: Parses and routes messages
- **Processor Layer**: Executes business logic
- **Store Layer**: Simulates transaction persistence

See: `/docs/architecture.md`

---

## 🔁 Transaction Flow

1. Client sends transaction request (0200)
2. Server parses message
3. Business logic validates transaction
4. Response is generated (0210)
5. Latency and result are logged

---

## ⚙️ Tech Stack

- Go (Golang)
- TCP Socket Programming
- JSON (for simplified message structure)
- Docker (optional)

---

## 📌 Notes

This project is for demonstration purposes only.  
Most real-world implementations involve secure ISO8583 parsing, HSM integration, and strict compliance (PCI-DSS).

---

## 👤 Author

Most of my production work is in private enterprise repositories, particularly in banking and payment systems.  
This repository is created to demonstrate architectural approach and system design thinking.
