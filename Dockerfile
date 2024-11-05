FROM golang:1.23-alpine

WORKDIR /app

COPY . .

RUN go mod download

RUN go mod tidy

RUN go build -o /sre_bootcamp

EXPOSE 4000

CMD [ "/sre_bootcamp" ]