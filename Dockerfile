FROM golang:1.26-alpine AS builder

WORKDIR /jam

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -trimpath -ldflags="-s -w" -o jam ./cmd/jam

FROM scratch

COPY --from=builder /jam/jam /jam/jam

WORKDIR /jam

CMD [ "./jam" ]
