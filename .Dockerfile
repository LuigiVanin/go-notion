FROM golang:1.21.3

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

EXPOSE 3000

RUN CGO_ENABLED=0 GOOS=linux go build -o go-notion

CMD ["go-notion"]

