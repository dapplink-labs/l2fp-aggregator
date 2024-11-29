# Build tss in a stock Go builder container
FROM golang:1.21-alpine3.18 as builder

RUN apk add --no-cache make gcc musl-dev linux-headers git

COPY . /app/

WORKDIR /app/manta-relayer
RUN make build

# Pull tss into a second stage deploy alpine container
FROM alpine:3.18

WORKDIR /app

RUN apk add --no-cache ca-certificates
COPY --from=builder /app/manta-relayer/mr /app

# EXPOSE 8545 8546 8547
ENTRYPOINT ["./mr"]
