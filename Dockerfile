# builder image
FROM golang:1.20.5

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . ./

RUN go build -o app ./cmd/server.go

EXPOSE 8080

CMD [ "./app","-p","8080" ]