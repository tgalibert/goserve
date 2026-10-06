# 🚀 GoServe

<p align="center">
  <strong>A lightweight, zero-dependency, production-grade HTTP/1.1 server built entirely from scratch in Go.</strong>
</p>

<p align="center">
  <a href="https://golang.org"><img src="https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat-square&logo=go" alt="Go Version"></a>
  <a href="https://www.rfc-editor.org/rfc/rfc9112"><img src="https://img.shields.io/badge/Spec-RFC%209112-blue?style=flat-square" alt="RFC 9112"></a>
  <a href="#"><img src="https://img.shields.io/badge/Dependencies-Zero-success?style=flat-square" alt="Zero Dependencies"></a>
  <a href="#"><img src="https://img.shields.io/badge/License-MIT-purple?style=flat-square" alt="License"></a>
</p>

---

## Overview

**GoServe** is a complete HTTP/1.1 server implementation built purely on raw TCP sockets provided by the Go standard package `net`. It deliberately avoids high-level abstractions like `net/http` to demonstrate how network I/O buffering, protocol grammar parsing, response serialization, and routing engines operate at the lowest levels.

Whether you're studying protocol engineering, designing embedded micro-services, or preparing a showcase project, **GoServe** provides an idiomatic, memory-safe, and clean foundation.

---

## Features

- **Zero External Dependencies** : Pure Go standard library (`net`, `io`, `bytes`, `strings`, `strconv`, `time`).
- **Stateful Buffered I/O (`IoParser`)** : Prevents body over-reading on continuous TCP streams while eliminating redundant syscalls.
- **RFC 9112 Compliant Request Parsing** :
- Strict Request Line validation (`METHOD`, `URI`, `PROTOCOL`).
- Safe query string extraction and path decoupling (`/items?page=1` $\rightarrow$ path + parameters map).
- Case-insensitive HTTP header normalization.
- Streaming body consumption using exact `Content-Length` accounting.
- **Fast Hierarchical Router** :
- Grouped by `(Path, Method)` for $O(1)$ lookup performance.
- Automatic `404 Not Found` dispatching.
- RFC 9110 compliant `405 Method Not Allowed` with automatic `Allow` header injection.
- **Modern CORS Engine & Preflight Support** :
- Transparent `OPTIONS` preflight caching with `Access-Control-Max-Age`.
- Dynamic multi-origin matching with `Vary: Origin` protection.
- **Ergonomic Response Builder** :
- Chainable helper methods (`SetJSON`, `SetHTML`, `SetHeader`).
- Automatic `Content-Length`, `Date` (RFC 1123 GMT), and `Connection: close` formatting.

---

## Architecture & Data Flow

```
                      +-----------------------------+
                      |   Client (Browser / cURL)   |
                      +-----------------------------+
                                     │  Raw TCP Connection
                                     ▼
                      +-----------------------------+
                      |     net.Listener (Accept)   |
                      |   Goroutine per Connection  |
                      +-----------------------------+
                                     │
                                     ▼
                      +-----------------------------+
                      |    IoParser (Stateful Buf)  |
                      |  - ReadLine() (CRLF slice)  |
                      |  - ReadBytes(N) (Body flux) |
                      +-----------------------------+
                                     │
                                     ▼
                      +-----------------------------+
                      |       Request Parser        |
                      |  - Method & Target & Proto  |
                      |  - Headers & Content-Length |
                      |  - Raw Body Bytes           |
                      +-----------------------------+
                                     │
                ┌────────────────────┴────────────────────┐
                ▼ (OPTIONS Preflight)                     ▼ (Standard Requests)
      +--------------------+                    +--------------------+
      | 204 No Content     |                    | Router & Handlers  |
      | CORS Allow Headers |                    | - 404 Not Found    |
      +--------------------+                    | - 405 + Allow      |
                │                               | - Custom Handlers  |
                │                               +--------------------+
                │                                         │
                └────────────────────┬────────────────────┘
                                     ▼
                      +-----------------------------+
                      |       Response Writer       |
                      |  HTTP/1.1 <Code> <Reason>   |
                      |  Headers + CRLF + Body      |
                      +-----------------------------+
                                     │  Raw Bytes
                                     ▼
                      +-----------------------------+
                      |        Socket net.Conn      |
                      +-----------------------------+
```

---

## 🚀 Quick Start

### 1. Installation

Clone the repository and inspect the codebase:

```bash
git clone https://github.com/your-username/goserve.git
cd goserve
```

### 2. Minimal Example

Create your `main.go` file:

```go
package main

import (
	"goserve/internal/network"
	"log"
)

func main() {
	// 1. Initialize server on port 8080
	server := network.NewServer(8080)

	// 2. Enable CORS with allowed origins (or leave empty for wildcard "*")
	server.AllowCors("http://localhost:3000", "https://myapp.com")

	// 3. Register routes
	server.Router.GET("/", func(req *network.Request, res *network.Response) {
		res.SetHTML([]byte("<h1>Welcome to GoServe!</h1>"))
	})

	server.Router.GET("/api/ping", func(req *network.Request, res *network.Response) {
		res.SetJSON([]byte(`{"message": "pong", "status": "ok"}`))
	})

	server.Router.POST("/api/echo", func(req *network.Request, res *network.Response) {
		// Echo client body with JSON content type
		res.SetJSON(req.Body)
	})

	// 4. Start listening
	log.Println("Server running on http://localhost:8080")
	if err := server.Start(); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
```

### 3. Run and Test

```bash
go run main.go
```

Test with `cURL`:

```bash
# Test GET request
curl -i http://localhost:8080/api/ping

# Test CORS Preflight
curl -i -X OPTIONS http://localhost:8080/api/ping \
  -H "Origin: http://localhost:3000" \
  -H "Access-Control-Request-Method: GET"

# Test 405 Method Not Allowed
curl -i -X DELETE http://localhost:8080/api/ping

# Test 404 Not Found
curl -i http://localhost:8080/unknown
```

---

## 📚 API Reference

### `Server`

| Method                                | Description                                                   |
| :------------------------------------ | :------------------------------------------------------------ |
| `NewServer(port int) *Server`         | Creates a new server instance with an initialized router.     |
| `server.Start() error`                | Binds the TCP listener and starts the concurrent accept loop. |
| `server.AllowCors(origins ...string)` | Enables CORS. Accepts specific origins or defaults to `*`.    |

### `Router`

Register handlers using standard HTTP verb methods:

```go
server.Router.GET(path string, handler HandleFunc)
server.Router.POST(path string, handler HandleFunc)
server.Router.PUT(path string, handler HandleFunc)
server.Router.DELETE(path string, handler HandleFunc)
server.Router.PATCH(path string, handler HandleFunc)
```

Where `HandleFunc` has the following signature:

```go
type HandleFunc func(req *network.Request, res *network.Response)
```

### `Request`

Inspect incoming HTTP request attributes:

```go
req.Method           // network.RequestMethod (e.g. GET, POST)
req.Path             // Clean route path (e.g. "/api/items")
req.QueryParameters  // map[string]string (e.g. ?page=2 -> {"page": "2"})
req.Headers          // Normalized map[string]string (lowercase keys)
req.ProtocolVersion  // network.RequestProtocol (e.g. "HTTP/1.1")
req.Body             // Raw byte slice ([]byte) of request payload
```

### `Response`

Configure output headers, status codes, and body payloads:

```go
// Set custom status code (typed enum)
res.StatusCode = constant.StatusCreated // 201

// Helper writers (automatically sets Content-Type & Content-Length)
res.SetJSON([]byte(`{"success": true}`))
res.SetHTML([]byte("<p>Hello</p>"))

// Raw body assignment
res.SetBody([]byte("plain data"))

// Custom headers
res.SetHeader("X-Custom-Header", "value")
```

---

## 🛡️ HTTP Specification Compliance

| Specification        | Feature             | Implementation                                                     |
| :------------------- | :------------------ | :----------------------------------------------------------------- |
| **RFC 9112 §3**      | Request Line Syntax | Strict token extraction (`METHOD SP Target SP Version CRLF`)       |
| **RFC 9112 §5**      | Header Parsing      | Delimited by CRLF, case-insensitive mapping, empty line terminator |
| **RFC 9112 §6**      | Content-Length Body | Byte-exact streaming via stateful buffer accumulation              |
| **RFC 9110 §15.5.6** | Method Not Allowed  | Status `405` with mandatory comma-separated `Allow` header         |
| **RFC 9110 §6.6.1**  | Date Header         | Automatic UTC timestamps formatted to RFC 1123 GMT                 |
| **W3C CORS**         | Preflight Options   | Automatic `204 No Content` response with cached policy headers     |

---

## 🗺️ Roadmap & Progress

- [x] **Phase 1**: TCP Network Socket & Concurrency Foundations
- [x] **Phase 2**: RFC 9112 Request Parser & Stateful I/O Buffer
- [x] **Phase 3**: HTTP/1.1 Response Generator & Serializer
- [x] **Phase 4**: Standard Headers & Dynamic Multi-Origin CORS
- [x] **Phase 5**: Hierarchical Routing Engine & RFC Error Dispatching
- [ ] **Phase 6**: Robustness & Performance (Deadlines/Timeouts, Keep-Alive, Panic Recovery)

---

## 📄 License

This project is open-source and available under the [MIT License](LICENSE).
