FROM golang:1.27 AS build

WORKDIR /tcp-chat

COPY go.mod ./


RUN go mod download


COPY . .

# create a static build CGO_ENABLED=0

RUN CGO_ENABLED=0 go build -o ./bin/main ./cmd


FROM alpine:3.24 

WORKDIR /tcp-chat


COPY --from=build ./tcp-chat/bin/main ./bin/main


EXPOSE 3000


CMD [ "./bin/main" ]



