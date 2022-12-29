FROM golang:1.16-alpine as builder

WORKDIR /app

COPY . .
RUN go mod download

COPY *.go ./

# Use CGO_ENABLED=0 to build a static binary
RUN CGO_ENABLED=0 go build -o /docker-go-redis


# Use the smallest possible image
# If you run `trivy image vue-go-redis-composant_go-app` ; you can see previous image has sev1 vulnerabilities
FROM scratch

COPY --from=builder /docker-go-redis docker-go-redis

EXPOSE 8080

CMD [ "/docker-go-redis" ]