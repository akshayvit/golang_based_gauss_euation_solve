FROM golang:1.16-alpine

WORKDIR /
COPY go.mod ./

RUN go mod download
COPY *.go ./
COPY *.go /gauss

RUN go build -o /docker-gauss .

EXPOSE 8080

USER noroot:noroot

CMD ["/docker-gauss"]