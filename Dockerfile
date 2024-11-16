FROM golang:1.22 AS builder

ENV CGO_ENABLED=0 GOOS=linux

WORKDIR /build

COPY . .

RUN go mod download

RUN go build -o rest-api ./cmd/rest-api/ && go build -o migrator ./cmd/migrator/

FROM alpine:latest AS production

WORKDIR /app

COPY --from=builder /build/rest-api .
COPY --from=builder /build/migrator .
COPY ./migrations/ ./migrations/

COPY ./entrypoint.sh .

RUN chmod +x ./entrypoint.sh && chmod +x ./entrypoint.sh && chmod +x ./migrator && chmod +x ./rest-api

RUN apk add --no-cache ca-certificates 

ENTRYPOINT ["./entrypoint.sh"]