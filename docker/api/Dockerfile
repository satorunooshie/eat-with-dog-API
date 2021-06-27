FROM golang:1.16

ENV GO111MODULE=on

WORKDIR /go/src/github.com/satorunooshie/eat-with-dog-API

COPY go.mod go.sum ./

RUN go mod download

COPY . .
RUN go build .

RUN go get github.com/pilu/fresh

EXPOSE 8888

CMD ["fresh"]
