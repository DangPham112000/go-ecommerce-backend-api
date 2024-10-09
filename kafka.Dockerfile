FROM golang:alpine AS builder 

WORKDIR /build 

COPY . .

RUN go mod download

RUN go build -o golangapp ./cmd/cli/kafka/main.kafka.go

FROM scratch

COPY --from=builder /build/golangapp / 

ENTRYPOINT [ "/golangapp" ]