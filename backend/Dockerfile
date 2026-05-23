FROM golang:1.22-alpine AS builder

RUN apk add --no-cache gcc musl-dev tesseract-ocr-dev leptonica-dev

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=1 go build -o main cmd/api/main.go

FROM alpine:latest

RUN apk add --no-cache tesseract-ocr leptonica

WORKDIR /app

COPY --from=builder /app/main .
COPY .env.example .env

EXPOSE 8080

CMD ["./main"]
