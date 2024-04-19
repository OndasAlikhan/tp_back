FROM golang:1.22.2

WORKDIR /app

RUN go install github.com/cosmtrek/air@v1.51.0 \
    && go install github.com/swaggo/swag/cmd/swag@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go mod tidy

ENTRYPOINT ["air", "-c", ".air.toml"]