# Go Proxy Server (Gin-based) for ClassiCube-PPC

A lightweight HTTP proxy server written in Go using the Gin framework. It supports file forwarding, release version tracking, and download statistics, enabling the download of ClassiCube-PPC resources on older systems (it can also be used for other purposes).

---

## 🚀 Features

-  Proxy file downloads via `?url=...`
-  Track active downloads and total bytes transferred
-  Serve latest release info from `release.json`


---
## Requirements
- Go 1.21 or newer
- Internet access for proxying external URLs

---
# Setup
```bash
git clone https://github.com/andreiixedev/Proxy-PPC.git
cd NewProxy
go mod init newproxy
go get github.com/gin-gonic/gin@latest

>> Run the Server
go run main.go
```

Server will start on http://localhost:5090

# Build Executable (Windows)
```bash
go build -o proxy.exe main.go
```
Then run:
```bash
./proxy.exe
```

# 🐳 Docker (Optional)
```bash
FROM golang:1.21-alpine
WORKDIR /app
COPY . .
RUN go build -o proxy .
EXPOSE 5090
CMD ["./proxy"]
```

Build and run:
```bash
docker build -t newproxy .
docker run -p 5090:5090 newproxy
```

# Credits
Built with [gin](https://github.com/gin-gonic/gin) and ❤️ by Andreiixe.