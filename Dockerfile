FROM golang:1.20 AS base

RUN apt update

WORKDIR /app/invoice

COPY go.mod go.sum ./

COPY . .

RUN make build/bin

FROM alpine as binary

COPY --from=base /app/invoice/bin/invoice /app/invoice

CMD ["/app/invoice"]
