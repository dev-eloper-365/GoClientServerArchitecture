# Go Client-Server Example 🚀

This document explains the implementation and usage of a simple **TCP-based Client-Server architecture in Go**, found in the `go_client_server` repository. It serves as a basic introduction to network communication using Go's `net` package.

---

## Project Structure

```
go_client_server/
├── client/
│   └── main.go         # TCP Client implementation
├── server/
│   └── main.go         # TCP Server implementation
└── README.md
```

---

## How It Works

### Server
- Listens on TCP port `:8080`.
- Accepts incoming client connections.
- Reads client messages and responds with a confirmation.

### Client
- Connects to `localhost:8080`.
- Sends a predefined message to the server.
- Receives and prints the server's response.

---

## Getting Started

### Prerequisites
- [Go](https://golang.org/dl/) installed (version 1.16 or higher recommended).

### Running the Server
```bash
cd server
go run main.go
```
Output:
```
Server is listening on port 8080...
```

### Running the Client
In a separate terminal:
```bash
cd client
go run main.go
```
Expected output:
```
Client connected to server at localhost:8080
Message sent to server: Hello from client!
Server response: Hello Client, I received your message.
```

---

## Customization Ideas

- Modify messages in `main.go` to suit your needs.
- Add concurrency support (goroutines) to handle multiple clients.
- Secure communication with TLS.
- Transmit structured data (e.g., JSON).

---

## Additional Resources

- [Go net package documentation](https://pkg.go.dev/net)
- [Effective Go - Concurrency](https://golang.org/doc/effective_go#concurrency)
- [TCP Server Examples](https://github.com/golang/go/wiki/TCPServer)

---

## License

This project is open-source and available under the [MIT License](LICENSE).

---

## Author

Created by [dev-eloper-365](https://github.com/dev-eloper-365) ❤️

