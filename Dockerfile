FROM golang:alpine AS builder 

WORKDIR /build 

COPY . .

RUN go mod download

RUN go build -o crm.shopgolang.com ./cmd/server

FROM scratch

COPY ./config /config 

COPY --from=builder /build/crm.shopgolang.com / 

ENTRYPOINT [ "/crm.shopgolang.com", "config/local.yaml" ]