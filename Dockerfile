FROM golang:1.22.4-alpine
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -ldflags="-w -s" -o job_test ./cmd/main.go
RUN chmod +x job_test
RUN go test -v ./test
CMD ["./job_test"]