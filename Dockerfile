FROM golang:1.22-alpine AS builder
ENV GOTOOLCHAIN=auto
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o numbers-service ./cmd/server

FROM gcr.io/distroless/static:nonroot
WORKDIR /app
COPY --from=builder /app/numbers-service /app/numbers-service
ENV HTTP_ADDR=:8080
EXPOSE 8080
ENTRYPOINT ["/app/numbers-service"]
