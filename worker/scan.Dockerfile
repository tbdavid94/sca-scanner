FROM golang:1.19-alpine

WORKDIR /home

COPY worker_scan.go .
COPY .env .
COPY go.* .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build worker_scan.go

#ENTRYPOINT ["/home/worker_scan"]