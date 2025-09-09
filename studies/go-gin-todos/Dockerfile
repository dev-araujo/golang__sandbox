# Build stage
FROM golang:1.24.4-alpine As builder
WORKDIR /app
COPY go.mod go.sum ./

RUN go mod download
COPY . .

RUN go build -o /app/main ./cmd/api



# Final stage

FROM scratch
WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 8080

ENTRYPOINT [ "/app/main" ]