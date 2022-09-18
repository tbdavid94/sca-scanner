FROM golang:1.19-alpine

WORKDIR /home

COPY worker_search.go .
COPY .env .
COPY ../go.* .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build worker_search.go

#EXPOSE 8081
#ENTRYPOINT ["/home/worker_search"]