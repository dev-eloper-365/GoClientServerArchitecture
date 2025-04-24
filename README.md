# Go Client-Server Example

This guide demonstrates a basic TCP client-server architecture in Go, designed to facilitate understanding of network communication using the Go `net` package. The project is minimal and intended for educational purposes.

## Overview

The application consists of a TCP server and a TCP client:

- The **server** listens on a TCP port, accepts incoming connections, and responds to messages.
- The **client** connects to the server, sends a message, and prints the server's response.

## Directory structure

```
go_client_server/
├── client/
│   └── main.go         # TCP client implementation
├── server/
│   └── main.go         # TCP server implementation
└── README.md
```

## Prerequisites

To run this project, ensure the following:

- Go (version 1.16 or later) is installed. Download from [golang.org/dl](https://golang.org/dl/).

## Getting started

### Run the server

1. Open a terminal.
2. Navigate to the `server` directory:
   ```bash
   cd server
   ```
3. Run the server:
   ```bash
   go run main.go
   ```
4. The terminal displays:
   ```
   Server is listening on port 8080...
   ```

### Run the client

1. Open a new terminal window.
2. Navigate to the `client` directory:
   ```bash
   cd client
   ```
3. Run the client:
   ```bash
   go run main.go
   ```
4. Expected output:
   ```
   Client connected to server at localhost:8080
   Message sent to server: Hello from client!
   Server response: Hello Client, I received your message.
   ```

## How it works

### Server behavior

- Listens for TCP connections on port `8080`.
- Accepts incoming connections from clients.
- Reads incoming data and replies with an acknowledgment.

### Client behavior

- Establishes a TCP connection to `localhost:8080`.
- Sends a predefined message to the server.
- Receives and displays the server’s response.

## Customization

You can extend or modify the project in the following ways:

- **Concurrent connections**: Use goroutines to support multiple client connections.
- **Structured communication**: Exchange JSON data between client and server.
- **Security enhancements**: Implement TLS encryption.
- **Interactive interface**: Add CLI arguments or prompts for dynamic messages.

## Learn more

- [net package documentation (Go)](https://pkg.go.dev/net)
- [Effective Go: Concurrency](https://golang.org/doc/effective_go#concurrency)
- [Go TCP Server example](https://github.com/golang/go/wiki/TCPServer)

## License

This project is licensed under the [MIT License](LICENSE).

## Author

Created by [dev-eloper-365](https://github.com/dev-eloper-365).

---

For feedback or contributions, feel free to open an issue or submit a pull request.

