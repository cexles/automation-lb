FROM golang:1.20-alpine AS build

RUN apk add --no-cache gcc musl-dev linux-headers git

WORKDIR /usr/src/app/

COPY . /usr/src/app/

RUN go mod download

RUN go build -o bin/link-lb

FROM alpine:latest

WORKDIR /cexles-chainlink-lb

COPY --from=build /usr/src/app/bin/link-lb /cexles-chainlink-lb/clb

COPY config.yaml /cexles-chainlink-lb/config.yaml

CMD ["./clb"]