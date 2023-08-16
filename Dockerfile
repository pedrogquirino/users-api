FROM golang:1.19

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

RUN go mod tidy

COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build cmd/main.go

EXPOSE 3000
CMD ["./main"]