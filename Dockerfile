FROM golang:1.19

WORKDIR /app

COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum
COPY ./main.go ./main.go

RUN go build -o app .

ENTRYPOINT [ "./app" ]
