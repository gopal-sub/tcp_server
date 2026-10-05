FROM golang:1.27



WORKDIR /tcp-chat

COPY go.mod ./


RUN go mod download


COPY . .

RUN go build -o ./bin/main ./cmd


EXPOSE 3000


CMD [ "./bin/main" ]



