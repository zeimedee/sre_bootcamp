FROM golang:1.23-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

RUN go mod tidy

COPY . .

RUN go build -o /sre_bootcamp ./cmd/api/main.go

EXPOSE 4000

CMD [ "/sre_bootcamp" ]