# builder image
FROM golang:1.20.5-alpine3.18 as builder

RUN mkdir /build

COPY . /build/

WORKDIR /build

RUN CGO_ENABLED=0 GOOS=linux go build -a -o app ./cmd/server.go


# generate clean, final image for end users
FROM alpine:3.18

COPY --from=builder /build/app .

# executable
ENTRYPOINT [ "./app", "-p", "8080" ]
