FROM golang:1.16-alpine

WORKDIR /app

COPY . .
RUN go mod download

COPY *.go ./

RUN go build -o /docker-go-redis

EXPOSE 8080

CMD [ "/docker-go-redis" ]