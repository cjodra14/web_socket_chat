FROM golang:latest

WORKDIR /real_time_chat

RUN GO111MODULE=on

COPY go.mod go.sum ./

RUN go mod download

COPY . . 

RUN go build -o main .

EXPOSE 8010:8010

CMD ["./main"]