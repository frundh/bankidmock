FROM golang:1.21.5-alpine as builder

WORKDIR /app
COPY . .
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /bankidmock ./cmd/bankidmock

FROM scratch

COPY --from=builder /bankidmock /bankidmock

ENTRYPOINT ["/bankidmock"]
