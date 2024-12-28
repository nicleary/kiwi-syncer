FROM golang:1.23

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /kiwi-syncer

EXPOSE 8080

CMD ["/kiwi-syncer"]