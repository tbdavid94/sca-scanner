FROM golang:1.19-alpine

WORKDIR /home

COPY worker_* .
COPY go.* .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build worker_scan.go
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build worker_search.go