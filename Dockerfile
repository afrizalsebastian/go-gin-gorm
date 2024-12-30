FROM golang:alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

EXPOSE 8000

RUN go build -o ./bin/go_gin_gorm ./cmd/server/main.go \
    && chmod +x ./bin/go_gin_gorm

ENTRYPOINT ["./bin/go_gin_gorm"]