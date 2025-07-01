# Go Backend Series (Part 2): Logging, HTTPS, and Scalable Server Architecture

Welcome to **Part 2** of the _Go Backend Series_!
In this part, we enhanced our backend by adding:

- **Zerolog** for structured, high-performance logging
- **Rotating file logs** using Lumberjack
- **Automatic HTTPS support** with fallback to HTTP
- **Modular server architecture** to support multiple engines (REST, GraphQL, WebSocket)
- **Route grouping** for cleaner and scalable API structure

These changes help make your Go API more production-ready, secure, and maintainable.

---

## 🔧 Tech Stack

- **Zerolog** – Fast, structured logging
- **Lumberjack** – Log file rotation
- **Self-signed certs** – For local HTTPS

---
