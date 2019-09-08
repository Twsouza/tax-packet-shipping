FROM golang:1.13.0-alpine3.10

WORKDIR /go/src/tax-packet-shipping
COPY . .

RUN go install

CMD ["tax-packet-shipping", "config.json"]