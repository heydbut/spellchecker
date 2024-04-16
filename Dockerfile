FROM golang:1.22 as builder

WORKDIR /app

# Copy go mod and sum files
COPY go.* ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -v -o spellchecker ./cmd/server

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/spellchecker .
COPY --from=builder /app/data/ data/
COPY --from=builder /app/templates/ templates/

CMD ["./spellchecker"]
