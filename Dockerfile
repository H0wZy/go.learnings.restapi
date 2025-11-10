ARG GO_VERSION=1.25.3-alpine
FROM golang:${GO_VERSION} as builder

RUN apk add --no-cache gcc musl-dev

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=1 GOOS=linux go build -o main .
LABEL authors="h0wzy"

# Create the final image, running the API and exposing in the 8080 port
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/db ./db

ARG PORT=8080
ENV PORT=${PORT}
EXPOSE ${PORT}

CMD ["./main"]