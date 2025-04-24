# 🏠 Go Client-Server Architecture

A simple and modular Go implementation demonstrating client-server architecture with TCP communication.

> ⚙️ _Note: This is a side project to practice Go networking and concurrency concepts._

---

## 📖 Overview

This repository contains:
- 🔌 A TCP server that handles multiple clients concurrently
- �️ A TCP client to send and receive messages
- 🥛 Use of Goroutines and channels for concurrency management

---

## 🚀 Usage

### Running the Server
```bash
go run server.go
```

### Running the Client
```bash
go run client.go
```

Connect multiple clients to the server to test concurrent messaging.

---

## 📁 Repository Structure

```
GoClientServerArchitecture/
├── client.go      # TCP client implementation
├── server.go      # TCP server implementation
├── README.md
```

---

## 📦 Prerequisites

- Go 1.16+
- TCP port availability (default: 8080)

---

## ⚠️ Disclaimer

This project is intended for educational purposes only.

---

## 📄 License

📝 MIT License

---

## 👨‍💻 Author

Maintained by [dev-eloper-365](https://github.com/dev-eloper-365)  
![GitHub Logo](https://img.icons8.com/ios-glyphs/30/github.png)
